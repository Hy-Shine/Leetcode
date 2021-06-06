/* Write your PL/SQL query statement below */

-- db type: Oracle
-- Level: Medium

SELECT 
CASE WHEN mod(id,2)=0 THEN id-1 
	WHEN id=(SELECT count(id) FROM seat) AND MOD(id,2)=1 THEN id 
	WHEN MOD(id,2)=1 THEN id+1 
END id,
student
FROM seat ORDER BY id;