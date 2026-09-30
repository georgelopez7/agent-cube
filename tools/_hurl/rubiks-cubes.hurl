# (Test): Create a new rubiks cube and capture its ID
POST {{host}}/api/v1/rubiks-cubes
Content-Type: application/json
{
    "llm": {
        "provider": "openai",
        "model": "gpt-4.0"
    }
}
HTTP 201
[Captures]
cube_id: jsonpath "$.cube.id"
[Asserts]
jsonpath "$.cube.id" exists
jsonpath "$.cube.llm.provider" == "openai"
jsonpath "$.cube.llm.model" == "gpt-4.0"
jsonpath "$.cube.status" == "created"

# (Test): Fetch the created rubiks cube and verify its data
GET {{host}}/api/v1/rubiks-cubes/{{cube_id}}
HTTP 200
[Asserts]
jsonpath "$.cube.id" == "{{cube_id}}"
jsonpath "$.cube.llm.provider" == "openai"
jsonpath "$.cube.llm.model" == "gpt-4.0"
jsonpath "$.cube.status" == "created"
jsonpath "$.cube.createdAt" exists
jsonpath "$.cube.updatedAt" exists

# (Test): Update the rubiks cube
PUT {{host}}/api/v1/rubiks-cubes/{{cube_id}}
Content-Type: application/json
{
    "llm": {
        "provider": "openai",
        "model": "gpt-4o"
    },
    "status": "in_progress",
    "cube": {
        "centers": {
            "F": { "stickers": { "F": "green" } },
            "B": { "stickers": { "B": "blue" } },
            "R": { "stickers": { "R": "red" } },
            "L": { "stickers": { "L": "orange" } },
            "U": { "stickers": { "U": "white" } },
            "D": { "stickers": { "D": "yellow" } }
        },
        "edges": {
            "UF": { "stickers": { "U": "white", "F": "green" } },
            "UR": { "stickers": { "U": "white", "R": "red" } },
            "UB": { "stickers": { "U": "white", "B": "blue" } },
            "UL": { "stickers": { "U": "white", "L": "orange" } },
            "FR": { "stickers": { "F": "green", "R": "red" } },
            "FL": { "stickers": { "F": "green", "L": "orange" } },
            "BR": { "stickers": { "B": "blue", "R": "red" } },
            "BL": { "stickers": { "B": "blue", "L": "orange" } },
            "DF": { "stickers": { "D": "yellow", "F": "green" } },
            "DR": { "stickers": { "D": "yellow", "R": "red" } },
            "DB": { "stickers": { "D": "yellow", "B": "blue" } },
            "DL": { "stickers": { "D": "yellow", "L": "orange" } }
        },
        "corners": {
            "UFR": { "stickers": { "U": "white", "F": "green", "R": "red" } },
            "URB": { "stickers": { "U": "white", "R": "red", "B": "blue" } },
            "UBL": { "stickers": { "U": "white", "B": "blue", "L": "orange" } },
            "ULF": { "stickers": { "U": "white", "L": "orange", "F": "green" } },
            "DFR": { "stickers": { "D": "yellow", "F": "green", "R": "red" } },
            "DRB": { "stickers": { "D": "yellow", "R": "red", "B": "blue" } },
            "DBL": { "stickers": { "D": "yellow", "B": "blue", "L": "orange" } },
            "DLF": { "stickers": { "D": "yellow", "L": "orange", "F": "green" } }
        },
        "rotations": []
    }
}
HTTP 200
[Asserts]
jsonpath "$.cube.id" == "{{cube_id}}"
jsonpath "$.cube.llm.model" == "gpt-4o"
jsonpath "$.cube.status" == "in_progress"

# (Test): Fetch the updated rubiks cube and verify the changes
GET {{host}}/api/v1/rubiks-cubes/{{cube_id}}
HTTP 200
[Asserts]
jsonpath "$.cube.id" == "{{cube_id}}"
jsonpath "$.cube.llm.provider" == "openai"
jsonpath "$.cube.llm.model" == "gpt-4o"
jsonpath "$.cube.status" == "in_progress"

# (Test): List rubiks cubes with the default limit
GET {{host}}/api/v1/rubiks-cubes
HTTP 200
[Asserts]
jsonpath "$.cubes" count > 0
jsonpath "$.cubes[*].id" contains "{{cube_id}}"

# (Test): Request an invalid rubiks cube ID — should return 400
GET {{host}}/api/v1/rubiks-cubes/not-an-id
HTTP 400

# (Test): Request a non-existent but valid rubiks cube ID — should return 404
GET {{host}}/api/v1/rubiks-cubes/000000000000000000000000
HTTP 404

# (Test): Update a non-existent but valid rubiks cube ID — should return 404
PUT {{host}}/api/v1/rubiks-cubes/000000000000000000000000
Content-Type: application/json
{
    "llm": {
        "provider": "openai",
        "model": "gpt-4o"
    },
    "status": "in_progress",
    "cube": {
        "centers": {},
        "edges": {},
        "corners": {},
        "rotations": []
    }
}
HTTP 404

# (Test): Apply a rotation to the rubiks cube
POST {{host}}/api/v1/rubiks-cubes/{{cube_id}}/rotations
Content-Type: application/json
{
    "rotation": "F"
}
HTTP 200
[Asserts]
jsonpath "$.cube.id" == "{{cube_id}}"
jsonpath "$.cube.cube.Rotations" count == 1
jsonpath "$.cube.cube.Rotations[0].rotation" == "F"
jsonpath "$.cube.cube.Rotations[0].from_scramble" == false

# (Test): Apply a counter-clockwise rotation to the rubiks cube
POST {{host}}/api/v1/rubiks-cubes/{{cube_id}}/rotations
Content-Type: application/json
{
    "rotation": "F'"
}
HTTP 200
[Asserts]
jsonpath "$.cube.id" == "{{cube_id}}"
jsonpath "$.cube.cube.Rotations" count == 2
jsonpath "$.cube.cube.Rotations[1].rotation" == "F'"
jsonpath "$.cube.cube.Rotations[1].from_scramble" == false

# (Test): Apply rotation with an invalid rubiks cube ID — should return 400
POST {{host}}/api/v1/rubiks-cubes/not-an-id/rotations
Content-Type: application/json
{
    "rotation": "F"
}
HTTP 400

# (Test): Apply rotation to a non-existent but valid rubiks cube ID — should return 404
POST {{host}}/api/v1/rubiks-cubes/000000000000000000000000/rotations
Content-Type: application/json
{
    "rotation": "F"
}
HTTP 404
