'use client';

import * as React from 'react';
import { Sidebar } from '@/components/layout/sidebar';
import { Topbar } from '@/components/layout/topbar';
import { Button } from '@/components/ui/button';
import { Modal } from '@/components/ui/modal';
import {
  ProcessesList,
  ProjectDetailsCard,
} from '@/components/features/projects/projects-panels';
import {
  ProcessHeader,
  ProcessGraphCard,
  ArtifactsCard,
  MessagesCard,
} from '@/components/features/processes/process-panels';
import { useToast } from '@/hooks/use-toast';

// ----- Mock Data (page-local; replace with stores/API later)
const mockUser = {
  id: '1',
  name: 'Dr. Sarah Chen',
  email: 'sarah.chen@university.edu',
  roles: ['ProjectLeader', 'Reviewer'],
};

type Proc = {
  id: string;
  name: string;
  stage: string;
  state: 'draft' | 'pending' | 'approved' | 'rejected';
};
type Proj = {
  id: string;
  name: string;
  description: string;
  theme: string;
  members: number;
  processes: Proc[];
};

const mockProjects: Proj[] = [
  {
    id: '1',
    name: 'Climate Change Impact Study',
    description:
      'Analyzing the effects of climate change on coastal ecosystems',
    theme: 'Environmental Science',
    members: 8,
    processes: [
      { id: '1', name: 'Ethics Review', stage: 'Under Review', state: 'pending' },
      { id: '2', name: 'Data Collection Approval', stage: 'Draft', state: 'draft' },
    ],
  },
  {
    id: '2',
    name: 'Neural Network Architecture',
    description:
      'Developing new approaches to deep learning optimization',
    theme: 'Computer Science',
    members: 5,
    processes: [
      { id: '3', name: 'Funding Approval', stage: 'Approved', state: 'approved' },
    ],
  },
];

export default function PrototypePage() {
  const toast = useToast();

  const [currentView, setCurrentView] = React.useState<
    'home' | 'project' | 'process' | 'profile' | 'settings'
  >('home');
  const [selectedProjectId, setSelectedProjectId] = React.useState<string | null>(null);
  const [selectedProcessId, setSelectedProcessId] = React.useState<string | null>(null);

  const [showCreateProject, setShowCreateProject] = React.useState(false);
  const [showCreateProcess, setShowCreateProcess] = React.useState(false);
  const [showApprovalModal, setShowApprovalModal] = React.useState(false);

  const [newProject, setNewProject] = React.useState({ name: '', description: '', theme: '' });
  const [newProcess, setNewProcess] = React.useState({ name: '', template: '' });

  const selectedProject = mockProjects.find((p) => p.id === selectedProjectId) || null;
  const selectedProcess =
    selectedProject?.processes.find((p) => p.id === selectedProcessId) || null;

  const goHome = () => setCurrentView('home');
  const openProject = (id: string) => {
    setSelectedProjectId(id);
    setCurrentView('project');
  };
  const openProcess = (projectId: string, processId: string) => {
    setSelectedProjectId(projectId);
    setSelectedProcessId(processId);
    setCurrentView('process');
  };

  // Actions
  const handleCreateProject = () => {
    if (!newProject.name || !newProject.description) return;
    toast.success('Project created successfully');
    setShowCreateProject(false);
    setNewProject({ name: '', description: '', theme: '' });
    openProject('1'); // demo nav like prototype
  };

  const handleStartProcess = () => setShowCreateProcess(true);

  const handleCreateProcess = () => {
    if (!newProcess.name || !newProcess.template) return;
    toast.success('Process started successfully');
    setShowCreateProcess(false);
    setNewProcess({ name: '', template: '' });
  };

  const handleApproval = (approved: boolean) => {
    toast.success(approved ? 'Process approved successfully' : 'Process rejected');
    setShowApprovalModal(false);
  };

  // Views
  function HomeView() {
    return (
      <div className="p-8">
        <div className="mx-auto max-w-4xl">
          <h1 className="mb-2 text-3xl font-bold text-[hsl(var(--fg))]">
            Welcome back, Dr. Chen
          </h1>
          <p className="mb-8 text-[hsl(var(--muted))]">
            Manage your research projects and approval processes
          </p>

          <div className="grid grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-3">
            <div className="radii-lg border border-base bg-[hsl(var(--surface))] p-6 shadow-soft">
              <h3 className="mb-2 font-semibold text-[hsl(var(--fg))]">Recent Projects</h3>
              <div className="space-y-2">
                {mockProjects.slice(0, 2).map((project) => (
                  <button
                    key={project.id}
                    onClick={() => openProject(project.id)}
                    className="block w-full text-left text-sm text-[hsl(var(--brand-600))] hover:text-[hsl(var(--brand-700))]"
                  >
                    {project.name}
                  </button>
                ))}
              </div>
            </div>

            <div className="radii-lg border border-base bg-[hsl(var(--surface))] p-6 shadow-soft">
              <h3 className="mb-2 font-semibold text-[hsl(var(--fg))]">Pending Approvals</h3>
              <div className="space-y-2">
                <div className="text-sm text-[hsl(var(--muted))]">Ethics Review</div>
                <div className="text-sm text-[hsl(var(--muted))]">Data Collection Approval</div>
              </div>
            </div>

            <div className="radii-lg border border-base bg-[hsl(var(--surface))] p-6 shadow-soft">
              <h3 className="mb-2 font-semibold text-[hsl(var(--fg))]">Quick Actions</h3>
              <div className="space-y-2">
                <Button
                  variant="ghost"
                  size="sm"
                  className="w-full justify-start"
                  onClick={() => setShowCreateProject(true)}
                >
                  New Project
                </Button>
              </div>
            </div>
          </div>
        </div>
      </div>
    );
  }

  function ProjectView() {
    if (!selectedProject) return <div className="p-8 text-[hsl(var(--fg))]">Project not found</div>;

    return (
      <div className="p-8">
        <div className="mx-auto max-w-6xl">
          <div className="mb-6 flex items-center justify-between">
            <div>
              <h1 className="mb-2 text-3xl font-bold text-[hsl(var(--fg))]">
                {selectedProject.name}
              </h1>
              <p className="text-[hsl(var(--muted))]">{selectedProject.description}</p>
            </div>
            <Button onClick={handleStartProcess}>Start Process</Button>
          </div>

          <div className="grid grid-cols-1 gap-6 lg:grid-cols-3">
            <div className="lg:col-span-2">
              <div className="radii-lg border border-base bg-[hsl(var(--surface))] shadow-soft">
                <div className="border-b border-base p-6">
                  <h2 className="text-xl font-semibold text-[hsl(var(--fg))]">Processes</h2>
                </div>
                <ProcessesList
                  processes={selectedProject.processes}
                  onOpenProcess={(pid: string) => openProcess(selectedProject.id, pid)}
                />
              </div>
            </div>

            <ProjectDetailsCard
              theme={selectedProject.theme}
              members={selectedProject.members}
              onManageMembers={() => toast.info('Manage Members (placeholder)')}
              onStartProcess={handleStartProcess}
            />
          </div>
        </div>
      </div>
    );
  }

  function ProcessView() {
    if (!selectedProject || !selectedProcess)
      return <div className="p-8 text-[hsl(var(--fg))]">Process not found</div>;

    const canApprove =
      selectedProcess.state === 'pending' && mockUser.roles.includes('Reviewer');

    return (
      <div className="p-8">
        <div className="mx-auto max-w-6xl">
          <ProcessHeader
            name={selectedProcess.name}
            state={selectedProcess.state}
            stage={selectedProcess.stage}
            showActions={canApprove}
            onApprove={() => setShowApprovalModal(true)}
            onReject={() => setShowApprovalModal(true)}
          />

          <div className="grid grid-cols-1 gap-6 lg:grid-cols-2">
            <div className="space-y-6">
              <ProcessGraphCard />
              <ArtifactsCard />
            </div>
            <div>
              <MessagesCard />
            </div>
          </div>
        </div>
      </div>
    );
  }

  const content =
    currentView === 'home' ? (
      <HomeView />
    ) : currentView === 'project' ? (
      <ProjectView />
    ) : currentView === 'process' ? (
      <ProcessView />
    ) : currentView === 'profile' ? (
      <div className="p-8 text-[hsl(var(--fg))]">Profile (placeholder)</div>
    ) : (
      <div className="p-8 text-[hsl(var(--fg))]">Settings (placeholder)</div>
    );

  return (
    <div className="flex h-screen bg-[hsl(var(--bg))]">
      <Sidebar
        projects={mockProjects.map((p) => ({
          id: p.id,
          name: p.name,
          processes: p.processes.map((pr) => ({ id: pr.id, name: pr.name })),
        }))}
        currentView={currentView}
        selectedProjectId={selectedProjectId}
        selectedProcessId={selectedProcessId}
        onGoHome={goHome}
        onSelectProject={openProject}
        onSelectProcess={openProcess}
        onNewProject={() => setShowCreateProject(true)}
      />

      <div className="flex flex-1 flex-col overflow-hidden">
        <Topbar
          userName={mockUser.name}
          onOpenProfile={() => setCurrentView('profile')}
          onOpenSettings={() => setCurrentView('settings')}
          onSignOut={() => toast.info('Signed out (demo)')}
        />
        <main className="flex-1 overflow-auto">{content}</main>
      </div>

      {/* Create Project */}
      <Modal
        isOpen={showCreateProject}
        onClose={() => setShowCreateProject(false)}
        title="Create New Project"
      >
        <div className="space-y-4">
          <div>
            <label className="mb-2 block text-sm font-medium text-[hsl(var(--fg))]">
              Project Name
            </label>
            <input
              type="text"
              className="w-full radii-md border border-base bg-[hsl(var(--surface))] px-3 py-2 text-[hsl(var(--fg))] focus:outline-none focus:ring-2 focus:ring-[hsl(var(--ring))]"
              value={newProject.name}
              onChange={(e) => setNewProject((s) => ({ ...s, name: e.target.value }))}
            />
          </div>

          <div>
            <label className="mb-2 block text-sm font-medium text-[hsl(var(--fg))]">
              Description
            </label>
            <textarea
              rows={3}
              className="w-full radii-md border border-base bg-[hsl(var(--surface))] px-3 py-2 text-[hsl(var(--fg))] focus:outline-none focus:ring-2 focus:ring-[hsl(var(--ring))]"
              value={newProject.description}
              onChange={(e) => setNewProject((s) => ({ ...s, description: e.target.value }))}
            />
          </div>

          <div>
            <label className="mb-2 block text-sm font-medium text-[hsl(var(--fg))]">
              Theme
            </label>
            <select
              className="w-full radii-md border border-base bg-[hsl(var(--surface))] px-3 py-2 text-[hsl(var(--fg))] focus:outline-none focus:ring-2 focus:ring-[hsl(var(--ring))]"
              value={newProject.theme}
              onChange={(e) => setNewProject((s) => ({ ...s, theme: e.target.value }))}
            >
              <option value="">Select a theme</option>
              <option value="Environmental Science">Environmental Science</option>
              <option value="Computer Science">Computer Science</option>
              <option value="Medicine">Medicine</option>
            </select>
          </div>

          <div className="flex space-x-3 pt-4">
            <Button onClick={handleCreateProject}>Create Project</Button>
            <Button variant="ghost" onClick={() => setShowCreateProject(false)}>
              Cancel
            </Button>
          </div>
        </div>
      </Modal>

      {/* Start Process */}
      <Modal
        isOpen={showCreateProcess}
        onClose={() => setShowCreateProcess(false)}
        title="Start New Process"
      >
        <div className="space-y-4">
          <div>
            <label className="mb-2 block text-sm font-medium text-[hsl(var(--fg))]">
              Process Name
            </label>
            <input
              type="text"
              className="w-full radii-md border border-base bg-[hsl(var(--surface))] px-3 py-2 text-[hsl(var(--fg))] focus:outline-none focus:ring-2 focus:ring-[hsl(var(--ring))]"
              value={newProcess.name}
              onChange={(e) => setNewProcess((s) => ({ ...s, name: e.target.value }))}
            />
          </div>

          <div>
            <label className="mb-2 block text-sm font-medium text-[hsl(var(--fg))]">
              Template
            </label>
            <select
              className="w-full radii-md border border-base bg-[hsl(var(--surface))] px-3 py-2 text-[hsl(var(--fg))] focus:outline-none focus:ring-2 focus:ring-[hsl(var(--ring))]"
              value={newProcess.template}
              onChange={(e) => setNewProcess((s) => ({ ...s, template: e.target.value }))}
            >
              <option value="">Select a template</option>
              <option value="ethics-review">Ethics Review</option>
              <option value="funding-approval">Funding Approval</option>
              <option value="data-collection">Data Collection Approval</option>
            </select>
          </div>

          <div className="flex space-x-3 pt-4">
            <Button onClick={handleCreateProcess}>Start Process</Button>
            <Button variant="ghost" onClick={() => setShowCreateProcess(false)}>
              Cancel
            </Button>
          </div>
        </div>
      </Modal>

      {/* Approve/Reject */}
      <Modal
        isOpen={showApprovalModal}
        onClose={() => setShowApprovalModal(false)}
        title="Submit Approval"
      >
        <div className="space-y-4">
          <p className="text-[hsl(var(--muted))]">
            Are you sure you want to approve this process?
          </p>

          <div>
            <label className="mb-2 block text-sm font-medium text-[hsl(var(--fg))]">
              Comments (optional)
            </label>
            <textarea
              rows={3}
              className="w-full radii-md border border-base bg-[hsl(var(--surface))] px-3 py-2 text-[hsl(var(--fg))] focus:outline-none focus:ring-2 focus:ring-[hsl(var(--ring))]"
              placeholder="Add any comments..."
            />
          </div>

          <div className="flex space-x-3 pt-4">
            <Button onClick={() => handleApproval(true)}>Approve</Button>
            <Button variant="destructive" onClick={() => handleApproval(false)}>
              Reject
            </Button>
            <Button variant="ghost" onClick={() => setShowApprovalModal(false)}>
              Cancel
            </Button>
          </div>
        </div>
      </Modal>
    </div>
  );
}
