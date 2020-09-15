### API Docs ###
# Create question

Create a new question

**URL** : `/questions`

**Method** : `POST`

**Auth required** : NO

**Data constraints**

```json
{
    "title": string,
    "description": string
}
```

**Data example**

```json
{
    "title": "Which movie has won more Oscars on a single edition?",
    "description": "I wanna know which movie has won the most Oscars on a signle edition."
}
```

## Success Response

**Code** : `201 Created`

**Content example**

```json
{
    "created_id": "5f60da76e19b01f994c039b2"
}
```

# Get All

Fetches all questions with its answers

**URL** : `/questions`

**Method** : `GET`

**Auth required** : NO

## Success Response

**Code** : `200 OK`

**Content example**

```json
{
    "questions": [
        {
            "id": "5f5ec2d19c69506a5bd27c8d",
            "title": "What is love?",
            "description": "I wanna know what love is",
            "answers": [
                {
                    "id": "5f5fa659d0cbd58658f6cf52",
                    "answer": "Baby don't hurt me"
                },
                {
                    "id": "5f5fa6b0866ae7f13c953342",
                    "answer": "Baby don't hurt me"
                },
                {
                    "id": "5f5fa7207d9281c3f9c46ba5",
                    "answer": "What is love is a song by Haddaway"
                }
            ]
        },
        {
            "id": "5f5ee1b6b3146856dd3a11b3",
            "title": "Best songs of The Smiths",
            "description": "I'm looking for the best song of The Smiths. I recently find out who they are and I've heard a couple of their songs but I wanna dig really hard into their best work",
            "answers": null
        },
        {
            "id": "5f60da76e19b01f994c039b2",
            "title": "Which movie has won more Oscars on a single edition?",
            "description": "I wanna know which movie has won the most Oscars on a signle edition.",
            "answers": null
        }
    ]
}
```

# Get Question by ID

Given a question id, fetches the matching question with its answers 

**URL** : `/questions/{question_id}`

**Method** : `GET`

**Auth required** : NO

## Success Response

**Code** : `200 OK`

**Content example**

```json
{
    "question": {
        "id": "5f60da76e19b01f994c039b2",
        "title": "Which movie has won more Oscars on a single edition?",
        "description": "I wanna know which movie has won the most Oscars on a signle edition.",
        "answers": null
    }
}
```

# Update Question

Given a question id, updates the matching question. Only the title and the description
are editable with this endpoint

**URL** : `/questions/{question_id}`

**Method** : `PATCH`

**Auth required** : NO

**Data constraints**

```json
{
    "title": string,
    "description": string
}
```

**Data example**

```json
{
    "title": "Which movie has won more Oscars ever"
}
```

## Success Response

**Code** : `200 OK`

**Content example**

```json
{
    "message": ""
}
```

# Delete Question

Given a question id, deletes the matching question

**URL** : `/questions/{question_id}`

**Method** : `DELETE`

**Auth required** : NO

## Success Response

**Code** : `200 OK`

**Content example**

```json
{
    "message": ""
}
```

# Create Answer

Given a question id, adds a new answer to that question

**URL** : `/questions/{question_id}/answers`

**Method** : `PUT`

**Auth required** : NO

**Data constraints**

```json
{
    "answer": string
}
```

**Data example**

```json
{
    "answer": "The Lord of the Rings: The Return of the King (2003)"
}
```

## Success Response

**Code** : `201 OK`

**Content example**

```json
{
    "created_id": "5f60da76e19b01f994c039c5"
}
```

# Update Answer

Create a new question

**URL** : `/questions/{question_id}/answers/{answer_id}`

**Method** : `PATCH`

**Auth required** : NO

**Data constraints**

```json
{
    "answer": string
}
```

**Data example**

```json
{
    "answer": "Ben-Hur (1959), Titanic (1997)  and The Lord of the Rings: The Return of the King (2003)"
}
```

## Success Response

**Code** : `200 OK`

**Content example**

```json
{
    "message": "success"
}
```
