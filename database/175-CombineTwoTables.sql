/* Write your PL/SQL query statement below */

-- db type: Oracle
-- Level: Easy

SELECT a.FirstName, a.LastName, b.City, b.State
FROM 
    Person a 
    LEFT JOIN Address b ON a.PersonId = b.PersonId

-- Write your MySQL query statement below
SELECT a.FirstName, a.LastName, b.City, b.State
FROM 
    Person a 
    LEFT JOIN Address b ON a.PersonId = b.PersonId