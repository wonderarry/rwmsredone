-- Accounts ---------------------------------------------------------------

CREATE TABLE accounts (
  id            TEXT PRIMARY KEY,
  first_name    TEXT,
  middle_name   TEXT,
  last_name     TEXT,
  grp           TEXT,                    -- GroupNumber
  created_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at    TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- Identities (many per account; local & external in one table) ----------

CREATE TABLE identities (
  id             TEXT PRIMARY KEY,
  account_id     TEXT NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
  provider       TEXT NOT NULL,          -- e.g. "local", "university-oidc"
  subject        TEXT NOT NULL,          -- unique per provider (login for local, sub for OIDC)
  email          TEXT,
  password_hash  TEXT,                   -- non-NULL only for local
  refresh_token  TEXT,                   -- optional (store only if you truly need it; encrypt at rest)
  expires_at     TIMESTAMPTZ,
  created_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
  UNIQUE (provider, subject)
);
CREATE INDEX idx_identities_account ON identities(account_id);
CREATE INDEX idx_identities_provider_subject ON identities(provider, subject);

-- Global roles -----------------------------------------------------------

CREATE TABLE account_global_roles (
  account_id TEXT NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
  role_key   TEXT NOT NULL,              -- matches domain.GlobalRole string
  PRIMARY KEY (account_id, role_key)
);

-- Projects ---------------------------------------------------------------

CREATE TABLE projects (
  id            TEXT PRIMARY KEY,
  name          TEXT NOT NULL,
  theme         TEXT,
  descr         TEXT,
  created_by    TEXT NOT NULL REFERENCES accounts(id),
  created_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at    TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- Project memberships (multirole allowed) --------------------------------

CREATE TABLE project_members (
  project_id  TEXT NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
  account_id  TEXT NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
  role_key    TEXT NOT NULL,             -- "ProjectLeader" | "ProjectMember"
  PRIMARY KEY (project_id, account_id, role_key)
);
CREATE INDEX idx_project_members_by_account ON project_members(account_id, role_key);
CREATE INDEX idx_project_members_by_project ON project_members(project_id, role_key);

-- Processes --------------------------------------------------------------

CREATE TABLE processes (
  id             TEXT PRIMARY KEY,
  project_id     TEXT NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
  template_key   TEXT NOT NULL,
  name           TEXT NOT NULL,
  current_stage  TEXT NOT NULL,
  state          TEXT NOT NULL CHECK (state IN ('active','completed','archived')),
  created_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at     TIMESTAMPTZ NOT NULL DEFAULT now()
);
CREATE INDEX idx_processes_by_project ON processes(project_id);
CREATE INDEX idx_processes_by_state ON processes(state);

-- Process memberships ----------------------------------------------------

CREATE TABLE process_members (
  process_id  TEXT NOT NULL REFERENCES processes(id) ON DELETE CASCADE,
  account_id  TEXT NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
  role_key    TEXT NOT NULL,             -- "Advisor" | "Student" | "Reviewer"
  PRIMARY KEY (process_id, account_id, role_key)
);
CREATE INDEX idx_process_members_by_account ON process_members(account_id, role_key);
CREATE INDEX idx_process_members_by_process ON process_members(process_id, role_key);

-- Approvals (idempotent upsert via PK) ----------------------------------

CREATE TABLE approvals (
  process_id     TEXT NOT NULL REFERENCES processes(id) ON DELETE CASCADE,
  stage_key      TEXT NOT NULL,
  by_account_id  TEXT NOT NULL REFERENCES accounts(id),
  by_role        TEXT NOT NULL,          -- matches domain.ProcessRole
  decision       TEXT NOT NULL CHECK (decision IN ('approve','reject')),
  comment        TEXT NOT NULL DEFAULT '',
  created_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
  PRIMARY KEY (process_id, stage_key, by_account_id)
);

-- Speeds up CountByDecisionAndRole(process_id, stage_key, by_role, decision)
CREATE INDEX idx_approvals_count ON approvals(process_id, stage_key, by_role, decision);

-- Outbox (transactional events) -----------------------------------------

CREATE TABLE outbox (
  id            BIGSERIAL PRIMARY KEY,
  topic         TEXT NOT NULL,
  payload       JSONB NOT NULL,
  created_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
  published_at  TIMESTAMPTZ
);
CREATE INDEX idx_outbox_unpublished ON outbox(published_at) WHERE published_at IS NULL;