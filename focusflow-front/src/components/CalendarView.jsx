import React from 'react';
import { Calendar, dateFnsLocalizer } from 'react-big-calendar';
import format from 'date-fns/format';
import parse from 'date-fns/parse';
import startOfWeek from 'date-fns/startOfWeek';
import getDay from 'date-fns/getDay';
import 'react-big-calendar/lib/css/react-big-calendar.css';

const locales = {
  'en-US': require('date-fns/locale/en-US'),
};

const localizer = dateFnsLocalizer({
  format,
  parse,
  startOfWeek,
  getDay,
  locales,
});

export default function CalendarView({ goals, selectedDate, setSelectedDate }) {
    if (!goals || !Array.isArray(goals)) {
      return <div>Loading calendar...</div>;
    }
  
    const events = [];
  
    goals.forEach(goal => {

      if (goal.createdAt) {
        events.push({
          title: `🎯 Goal: ${goal.title}`,
          start: new Date(goal.createdAt),
          end: new Date(goal.createdAt),
          allDay: true,
        });
      }

      if (Array.isArray(goal.tasks)) {
        goal.tasks.forEach(task => {
          if (task.createdAt) {
            events.push({
              title: `📝 ${task.title} (${goal.title})`,
              start: new Date(task.createdAt),
              end: new Date(task.createdAt),
              allDay: false,
            });
          }
        });
      }
    });
  
    const onSelectSlot = (slotInfo) => {
      setSelectedDate(slotInfo.start);
    };
  
    return (
        <div className="mb-4">
          <Calendar
            localizer={localizer}
            events={events}
            startAccessor="start"
            endAccessor="end"
            selectable
            style={{ height: 500 }}
            onSelectSlot={onSelectSlot}
            popup
            dayPropGetter={(date) => {
              const isSelected =
                date.getFullYear() === selectedDate.getFullYear() &&
                date.getMonth() === selectedDate.getMonth() &&
                date.getDate() === selectedDate.getDate();
              if (isSelected) {
                return {
                  style: {
                    backgroundColor: '#cce5ff',
                  },
                };
              }
              return {};
            }}
          />
        </div>
      );
    }