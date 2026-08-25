from dotenv import load_dotenv
import os
import tempfile
from pathlib import Path
from langchain_community.document_loaders import TextLoader
load_dotenv()

def load_text_file():
    # Create a temporary text file with sample content
    with tempfile.NamedTemporaryFile(delete=False, suffix=".txt") as temp_file:
        temp_file.write(b"This is a sample text file for testing.")
        temp_file_path = temp_file.name

    try:
        # Load the text file using TextLoader
        loader = TextLoader(temp_file_path)
        documents = loader.load()

        print(f"Loaded {len(documents)} document(s)")
        print(f"Content preview: {documents[0].page_content[:100]}")
        print(f"Metadata: {documents[0].metadata}")

        # for doc in documents:
        #     print(doc)
        #     print(f"Loaded document content: {doc.page_content}")

    finally:
        # Clean up the temporary file
        os.remove(temp_file_path)

    return documents

result = load_text_file()
print(result)