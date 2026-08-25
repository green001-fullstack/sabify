import chromadb

client = chromadb.PersistentClient(path="./chroma_db")

name = "business"

collection = client.get_or_create_collection(name)

collection.add(
    documents=[
        "Customers who receive personalized product recommendations are more likely to make repeat purchases.",
        "Small businesses often use cloud-based accounting software to manage expenses, invoices, and financial reports.",
        "Customer churn can be reduced by identifying customers with declining engagement and offering targeted support.",
        "Sales teams can improve conversion rates by analyzing customer behavior and prioritizing high-value leads."
    ],

    ids=[
        "business_doc_1",
        "business_doc_2",
        "business_doc_3",
        "business_doc_4"
    ],

    metadatas=[
        {
            "category": "customer_engagement",
            "department": "marketing",
            "year": 2026
        },
        {
            "category": "financial_management",
            "department": "finance",
            "year": 2026
        },
        {
            "category": "customer_retention",
            "department": "customer_success",
            "year": 2026
        },
        {
            "category": "sales",
            "department": "sales",
            "year": 2026
        }
    ]
)

data = collection.get()
print(data)
