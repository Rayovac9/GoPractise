## Go Advanced Boot Camp Week 2 Assignment

## Question

 When a sql.ErrNoRows is encountered at the dao level, should the error be wrapped and thrown to the higher level?

## Idea

Operations in the dao are object-specific. sql.ErrNoRows is a generic statement, not an expression.
For example:
ErrNoRows is a generic term, not an expression. For example, if you are looking for a user by username, and you can't find the user, you should wrap sql.ErrNoRows to a ‘user does not exist’ error.
This way, if there is some kind of generic processing in the upper layers, it can be accommodated.

## Code file description

### init.go

1. Prepare the test table and data (Sqlite). 2.

2. Initialise Conn

### dao.go

1. Manipulate the User's Dao

2. FindByName,QueryBySex

3. sql.ErrNoRows Wrap usage

### main.go

1. Examples of usage

## Compile

`go run github.com/webmin7761/go-school/homework/error`

Translated with DeepL.com (free version)
