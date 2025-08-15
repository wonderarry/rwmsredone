export type ValidationResult<TFields extends string> = {
  ok: boolean;
  errors: Partial<Record<TFields, string>>;
};

export function validateLogin(login: string, password: string): ValidationResult<'login' | 'password'> {
  const errors: Record<string, string> = {};
  if (!login?.trim()) errors.login = 'Login is required';
  if (!password?.trim()) errors.password = 'Password is required';
  return { ok: Object.keys(errors).length === 0, errors };
}

export function validateCreateProject(
  name: string,
  description: string,
  theme: string
): ValidationResult<'name' | 'description' | 'theme'> {
  const errors: Record<string, string> = {};
  if (!name?.trim()) errors.name = 'Name is required';
  if (!description?.trim()) errors.description = 'Description is required';
  if (!theme?.trim()) errors.theme = 'Theme is required';
  return { ok: Object.keys(errors).length === 0, errors };
}

export function validateCreateProcess(
  name: string,
  templateKey: string
): ValidationResult<'name' | 'templateKey'> {
  const errors: Record<string, string> = {};
  if (!name?.trim()) errors.name = 'Process name is required';
  if (!templateKey?.trim()) errors.templateKey = 'Template is required';
  return { ok: Object.keys(errors).length === 0, errors };
}