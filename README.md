# AI Learning Intelligence LMS

An AI-powered Learning Management System that helps teachers understand student performance and helps students learn more effectively.

> Status: MVP / Hackathon build
> Product type: AI-Powered Learning Management System
> Primary users: Teachers & Students

---

## Overview

Unlike a traditional LMS that mainly manages courses, materials, assignments, and grades, this platform uses AI to analyze how students are learning, identify knowledge gaps, recommend interventions, and connect students with suitable study partners.

For the MVP, the platform focuses on:

- Core LMS functionality
- AI Quiz Generator
- AI Teacher Assistant
- AI Learning Coach
- AI Study Group Matcher

To keep the demo fast and reliable, **objective/MCQ quizzes** are the primary assessment method in the MVP.

## Problem Statement

Teachers often have to manually create assessments, grade students, analyze performance, and identify which topics students are struggling with. Students, meanwhile, often receive a score without understanding:

- Which topics they are weak in
- What they should study next
- Which classmates could help them

The platform closes this loop:

```
Teach → Assess → Analyze → Personalize → Collaborate → Improve
```

## Core Vision

An LMS that doesn't just record grades — it understands student performance and uses AI to help students and teachers improve.

---

## MVP Goals

**A teacher can:**
- Create a course
- Upload learning materials
- Generate objective quizzes using AI
- Review and publish quizzes
- View student results
- See AI-generated class insights and common misconceptions
- Receive teaching recommendations

**A student can:**
- Enroll in courses
- Access learning materials
- Take objective quizzes and get instant grades
- See their strengths and weaknesses
- Interact with an AI Learning Coach
- Receive personalized practice recommendations
- Get matched with suitable study partners

---

## User Roles

### Teacher
Register/login, create and manage courses, upload materials, generate quizzes, edit/review AI-generated questions, publish quizzes, view student results, view AI class insights and student performance.

### Student
Register/login, browse/enroll in courses, view materials, take quizzes, view results and learning progress, use the AI Learning Coach, create/update a learning profile, view study group recommendations.

---

## Feature Breakdown

### 1. Core LMS
- **Authentication:** register, login, logout, role-based access, protected dashboards
- **Courses:** create/edit/delete courses, descriptions, modules/topics, learning materials
- **Enrollment:** browse, enroll, and access enrolled course materials

### 2. AI Quiz Generator
Teachers specify course, topic, number of questions, difficulty, and question type. The AI generates each question with four options, the correct answer, topic tag, difficulty, and an explanation. Teachers can edit, delete, regenerate, or change answers/difficulty before publishing.

### 3. Student Quiz System
Students open available quizzes, answer and navigate questions, submit, and receive instant results (score, correct/incorrect breakdown), with options to view performance or ask the AI Coach.

### 4. AI Teacher Assistant
Analyzes quiz results across the class:
- **Class performance:** average, highest/lowest topics, completions, distribution
- **Topic analysis:** performance broken down by topic (each question is tagged)
- **Misconception detection:** AI identifies patterns in incorrect answers
- **Teaching recommendations:** suggested topic to revisit, explanation, teaching approach, practice questions, and a short intervention lesson

### 5. AI Learning Coach
A personalized student assistant that uses course materials, quiz results, weak topics, and learning progress — not a generic chatbot.
- Personalized recommendations after each quiz
- AI tutor for explanations, examples, and practice questions (via RAG over course materials where available)
- Personalized practice question generation for weak topics, which updates the learning profile

### 6. AI Study Group Matcher
Builds a **Student Learning Profile** (enrolled courses, strong/weak topics, availability, preferred study format, optional location) and matches students based on:
- Same course
- Complementary strengths and weaknesses
- Similar learning goals
- Compatible availability and study format

### 7. Study Group
Matched students can view group members, shared course, recommended focus topics, and use a basic group chat/study space.

### 8. Dashboards
- **Student dashboard:** enrolled courses, progress, upcoming assignments, recent grades, weak topics, AI recommendations, study group matches
- **Teacher dashboard:** total courses/students, active quizzes, average class performance, recent submissions, AI-generated insights, students needing attention

### 9. Roadmap
Here, when a student logs in to his dashboard, he will be asked questions on what he wants to learn and the AI will generate a roadmap journey based on his answer

---

## Architecture

```
Web Application
      ↓
Next.js Frontend
      ↓
Backend API
      ↓
PostgreSQL / AI Service / File Storage
      ↓
LLM + RAG
```

### Recommended Stack

| Layer          | Technologies |
|----------------|--------------|
| Frontend       | Next.js, React, TypeScript, Tailwind CSS |
| Backend        | NestJS, TypeScript |
| AI             | Python, FastAPI, LLM API, Embeddings, RAG |
| Database       | PostgreSQL, pgvector |
| Authentication | JWT |

### Core Database Entities

`User`, `Course`, `Enrollment`, `CourseMaterial`, `Quiz`, `Question`, `QuizAttempt`, `QuizAnswer`, `LearningProfile`, `LearningProgress`, `AIInsight`, `StudyGroup`, `StudyGroupMember`

Each quiz question stores its topic and difficulty, enabling topic-level performance calculations.

---

## Demo Flow (~3–5 minutes)

1. **Teacher** creates "Physics 101" and uploads "Newton's Laws" materials.
2. **AI Quiz Generation** — teacher requests 10 MCQs; AI generates the quiz; teacher reviews and publishes.
3. **Students** — two demo students take the quiz (e.g., 4/10 and 8/10).
4. **AI Analysis** — teacher dashboard shows class average, weakest topic, and the share of students who struggled.
5. **AI Teacher Assistant** — teacher asks "What should I do about this?" and receives an intervention plan.
6. **Student** — the lower-scoring student opens the AI Learning Coach and receives personalized practice questions.
7. **Study Group Matcher** — the platform matches the two students based on complementary strengths and weaknesses.

This flow demonstrates the full product concept without requiring long-form written answers.

---

## MVP Scope

### Must Have
- Authentication with teacher/student roles
- Course creation, enrollment, and materials
- AI quiz generation and MCQ quiz system
- Automatic grading and topic-level performance
- AI Teacher Assistant and class insights
- AI Learning Coach
- Student learning profile
- AI Study Group Matcher with basic study group page/chat
- Responsive dashboards

### Future
- AI essay/theory grading
- Assignment understanding verification
- AI Career Navigator
- Academic Second Brain
- AI group tutor
- Institution-level analytics
- Mobile application
- Advanced communication / video classes

---

## Success Criteria

The MVP is successful if a user can complete this end-to-end flow:

```
Create course → Generate AI quiz → Students take quiz → Automatic grading
→ AI identifies learning gaps → Teacher receives recommendations
→ Students receive personalized AI support → AI matches complementary
  students into a study group
```

**Value delivered:**
- **Teachers** — reduce assessment workload and understand where students are struggling.
- **Students** — understand their weaknesses, receive personalized help, and find peers who complement their learning.

---

## Status

This README is derived from the product requirements document (v1.1) and reflects the planned MVP scope. Implementation details (setup, install instructions, environment variables, scripts) should be added here as the codebase comes online.