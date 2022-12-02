# 📖 **Mini Blog Backend**

- Using MySQL
- Using GORM as ORM
- Using Chi as router

## 🚀 **How to Run**

> `go run ./cmd/server`

## 📑 **Database schema**

```
CREATE TABLE `blog_post` (
  `id` INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  `title` VARCHAR(255) NOT NULL,
  `content` LONGTEXT NOT NULL,
);
```

## 🔖 **Sample Data**

> blog_post table

| id  | title                | content                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| --- | -------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 1   | What is Lorem Ipsum? | Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum. |

## ☕ **Endpoint List**

### **1. GET /posts**

This endpoint will return list of blog post.

#### **Request**

> `GET /posts`

#### **Response**

Status code: `200`

Body:

```
[
  {
    "id": 1,
    "title": "What is Lorem Ipsum?",
    "content": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."
  }
]
```

### **2. POST /posts**

This endpoint will add new blog post to database.

#### **Request**

> `POST /posts`

Body:

```
{
  "title": "What is Git?",
  "content": "Git is a free and open source distributed version control system designed to handle everything from small to very large projects with speed and efficiency."
}
```

### **Response**

Status code: `200`

Body:

```
{
  "id": 2,
  "title": "What is Git?",
  "content": "Git is a free and open source distributed version control system designed to handle everything from small to very large projects with speed and efficiency."
}
```
