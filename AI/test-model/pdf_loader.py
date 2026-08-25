from dotenv import load_dotenv
import os
import tempfile
from pathlib import Path
from langchain_community.document_loaders import PyPDFLoader
load_dotenv()

def pdf_loader(pdf_path: str):
    loader = PyPDFLoader(pdf_path)
    documents = loader.load()

    print(f"Loaded {len(documents)} documents from PDF")
    for i, doc in enumerate(documents):
        print(f"Document {i+1} content preview: {doc.page_content[:100]}")
        print(f"Metadata: {doc.metadata}")


result = pdf_loader("LMS_AI_Architecture_and_Coding_Plan.pdf")
print(result)
