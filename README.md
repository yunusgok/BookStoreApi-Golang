# HW1- Library

This project is a basic library application for listing and searching books

## Usage
### List
To list all books in the library
```bash
go run main.go list
```
Output
```b
A Bend in the River
A Brief History of Seven Killings
A Christmas Carol
A Clockwork Orange
A Confederacy of Dunces
A Dance to the Music of Time
A Death in the Family
.
.
.
```
### Search
To search books with specific name
```bash
go run main.go search <Words to be searched>
```
Example
```bash
go run main.go search harry
```
Output
```
Harry Potter And The Chamber Of Secrets
Harry Potter And The Philosopher's Stone
Harry Potter And The Prisoner Of Azkaban
Harry Potter and the Deathly Hallows
Harry Potter and the Goblet of Fire
Harry Potter and the Half-Blood Prince
Harry Potter and the Order of the Phoenix
```

## License
[MIT](https://choosealicense.com/licenses/mit/)