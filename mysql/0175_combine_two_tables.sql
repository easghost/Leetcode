-- Runtime: 317 ms
-- Memory Usage: 0 MB

SELECT FirstName, LastName, City, State 
FROM Person 
LEFT JOIN Address 
ON Person.PersonId = Address.PersonId