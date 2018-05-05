import json
import jsonschema
 
schema = open("schema.json").read()
print schema
 
data = open("data.json").read()
print data
 
try:
    jsonschema.validate(json.loads(data), json.loads(schema))
except jsonschema.ValidationError as e:
    print e.message
except jsonschema.SchemaError as e:
    print e

try:
    v = jsonschema.Draft3Validator(json.loads(schema))
    for error in sorted(v.iter_errors(json.loads(data)), key=str):
        print(error.message)
except jsonschema.ValidationError as e:
    print e.message
