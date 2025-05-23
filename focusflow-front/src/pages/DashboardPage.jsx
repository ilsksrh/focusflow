import React, { useEffect, useState } from 'react';
import Sidebar from '../components/Sidebar';
import WelcomeBanner from '../components/WelcomeBanner';
import AdminUserPanel from '../components/AdminUserPanel';
import CalendarView from '../components/CalendarView';
import TaskList from '../components/TaskList';
import GoalForm from '../components/GoalForm';
import ProfilePanel from '../components/ProfilePanel';
import { getCurrentUser } from '../api/auth';
import { getGoals, createGoal } from '../api/goals';
import { createTask } from '../api/tasks';
import { format } from 'date-fns';

export default function DashboardPage() {
  const [user, setUser] = useState(null);
  const [loading, setLoading] = useState(true);
  const [viewMode, setViewMode] = useState('calendar');
  const [goals, setGoals] = useState([]);
  const [selectedDate, setSelectedDate] = useState(new Date());
  const [showGoalForm, setShowGoalForm] = useState(false);
  const [showProfile, setShowProfile] = useState(false);

  useEffect(() => {
    async function fetchData() {
      try {
        const [userData, goalsData] = await Promise.all([
          getCurrentUser(),
          getGoals(),
        ]);

        // ❌ Если юзер заблокирован — редиректим
        if (userData.is_blocked) {
          alert('Your account is blocked. Please contact support.');
          window.location.href = '/';
          return;
        }

        setUser(userData);
        setGoals(goalsData);
      } catch (err) {
        window.location.href = '/'; // если не авторизован
      } finally {
        setLoading(false);
      }
    }

    fetchData();
  }, []);

  if (loading) return null; // ничего не рендерим, пока грузится

  const handleCreateGoal = async (goalData) => {
    const newGoal = await createGoal(goalData);
    setGoals(prev => [...prev, newGoal]);
    setShowGoalForm(false);
  };

  const handleCreateTask = async (taskData) => {
    const newTask = await createTask(taskData);
    setGoals(prevGoals => prevGoals.map(goal =>
      goal.id === newTask.goal_id
        ? { ...goal, tasks: [...goal.tasks, newTask] }
        : goal
    ));
  };

  return (
    <div className="container-fluid">
      <div className="row">
        {/* Sidebar */}
        <div className="col-md-2 bg-light vh-100">
          <Sidebar
            viewMode={viewMode}
            setViewMode={setViewMode}
            onProfileClick={() => setShowProfile(true)}
            setShowProfile={setShowProfile}
            user={user}
          />
        </div>

        {/* Main content */}
        <div className="col-md-10">
          <div className="p-4">
            <WelcomeBanner user={user} />

            {showProfile ? (
              <ProfilePanel user={user} onUpdated={setUser} />
            ) : (
              <>
                {viewMode === 'calendar' && (
                  <>
                    <CalendarView
                      goals={goals}
                      selectedDate={selectedDate}
                      setSelectedDate={setSelectedDate}
                    />
                    <div className="mt-4">
                      <div className="d-flex align-items-center justify-content-between mb-3">
                        <h4>Manage Goals for {format(selectedDate, 'dd MMMM yyyy')}</h4>
                        <button
                          className="btn btn-primary btn-sm"
                          onClick={() => setShowGoalForm(prev => !prev)}
                        >
                          {showGoalForm ? 'Cancel' : 'Add Goal'}
                        </button>
                      </div>

                      {showGoalForm && (
                        <div className="mb-4">
                          <GoalForm onCreate={handleCreateGoal} selectedDate={selectedDate} />
                        </div>
                      )}

                      <TaskList
                        goals={goals}
                        selectedDate={selectedDate}
                        setGoals={setGoals}
                      />
                    </div>
                  </>
                )}

                {viewMode === 'kanban' && <div>Kanban view coming soon...</div>}
                {viewMode === 'list' && <div>List view coming soon...</div>}
                {viewMode === 'admin' && user?.role === 'admin' && <AdminUserPanel />}
              </>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}
