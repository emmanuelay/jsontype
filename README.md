# jsontype
Helper package for more granular control over deserialization of json payloads.

The sole purpose of this package was to be able to deserialize JSON payloads and effectively use them for updating a database.

Basically turning a struct like this:

```
	type User struct {
		FirstName 	jsontype.String `json:"first_name"`
		LastName 	jsontype.String `json:"last_name"`
		City 		jsontype.String `json:"city"`
		Country 	jsontype.String `json:"country"`
		Age 		jsontype.Int 	`json:"age"`
		PhoneNum	jsontype.String `json:"phone"`
		IsActive 	jsontype.Bool 	`json:"active"`
	}
```

..with a payload like this:

```
	[PUT] /1.0/user/1234
	{
		"first_name": "John",
		"last_name": "Doe",
		"age": 33,
		"active": true
	}
```

..into something like this:

```
	UPDATE 
		tbl_user 
	SET 
		first_name = 'John', 
		last_name = 'Doe', 
		age = 33, 
		active = true 
	WHERE 
		id = 1234
```

..by safely removing any missing or invalid values in the payload and only updating the values that are present in the payload, using the built-in **UnmarshalJSON** method.

# Credits
Based on an article by Jon Calhoun
https://www.calhoun.io/how-to-determine-if-a-json-key-has-been-set-to-null-or-not-provided/

I have simply added tests and support for string, int, int64 and bool as that is just what I needed.
I will probably add support for **time.Time** aswell.

# Other
This whole package probably can and will be replaced with generics in the near future.