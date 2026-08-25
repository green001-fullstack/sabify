import chromadb

client = chromadb.PersistentClient(path= "./chroma_db")

name = "vehicles"

collection = client.get_or_create_collection(name)

print(collection.name)

collection.add(
    documents=[
        "A car is a road vehicle, typically with four wheels, powered by an internal combustion engine or electric motor.",
        "A plane is an aircraft designed to fly through the air using wings and engines.",
        "A ship is a large watercraft designed to travel across oceans, seas, or other bodies of water.",
        "A bus is a large road vehicle designed to carry many passengers, usually along a fixed route."
    ],
    ids=[
        "car",
        "plane",
        "ship",
        "bus"
    ]
)

results = collection.query(
    query_texts=["Which is one travels through the ocean, ferrari, boat and bus"],
    n_results=1
)

print(results)


data = collection.get()
for i, doc in zip(data["ids"], data["documents"]):
    print(f"{i} -> {doc}")