from langchain_community.document_loaders import PyPDFLoader
from langchain_text_splitters import RecursiveCharacterTextSplitter

def loadpdf(pdf: str):
    loader = PyPDFLoader(pdf)
    documents = loader.load()
    return documents

def chunk_text(documents: list):
    text_splitter = RecursiveCharacterTextSplitter(
        chunk_size=800,
        chunk_overlap=100
    )

    chunks = text_splitter.split_documents(documents)
    return chunks

path = "./ss1_biology_first_term.pdf"
load = loadpdf(path)
chunks = chunk_text(load)

# for i, chunk in enumerate(chunks[:3]):
#     print(f"\nChunk {i + 1}:")
#     print(chunk.page_content)
#     print(chunk.metadata)

# print(len(chunks))
# print(chunks[0])

print("END OF CHUNK 1:")
print(chunks[0].page_content[-100:])

print("\nSTART OF CHUNK 2:")
print(chunks[1].page_content[:100])

for i in range(5):
    print(f"\nChunk {i + 1}: {len(chunks[i].page_content)} characters")

for i in range(4):
    current = chunks[i].page_content
    next_chunk = chunks[i + 1].page_content

    overlap = set()

    for length in range(1, min(len(current), len(next_chunk), 100) + 1):
        if current[-length:] == next_chunk[:length]:
            overlap.add(length)

    print(f"Chunk {i + 1} → Chunk {i + 2}: {max(overlap, default=0)} characters")