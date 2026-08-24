from pydantic import BaseModel

class Question(BaseModel):
    question: str
    options: dict[str, str]
    student_answer: str
    correct_answer: str


class AnalyzeRequest(BaseModel):
    student_id: str
    subject: str
    questions: list[Question]