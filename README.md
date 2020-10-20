# sqlvalidator
Validates SQL against different dialects.

Experiment using the Golang antlr4 target.

# Examples

```
sqlvalidator validate --dialect hive "SELECT * FROM table"
sqlvalidator validate --dialect mysql "SELECT * FROM USERS"
```
