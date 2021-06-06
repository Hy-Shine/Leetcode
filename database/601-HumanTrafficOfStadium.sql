/* Write your PL/SQL query statement below */

-- db type: Oracle
-- Level: Hard

SELECT t.id,to_char(t.visit_date,'yyyy-mm-dd') visit_date,t.people FROM (
SELECT id,visit_date,people,DENSE_RANK() OVER(ORDER BY id) r
FROM Stadium WHERE people>=100) t 
WHERE t.id-t.r IN (
	SELECT t2.id-t2.x FROM (
		SELECT id,DENSE_RANK() OVER(ORDER BY id) x 
		FROM Stadium WHERE people>=100) t2 
		GROUP BY (t2.id-t2.x) HAVING count(*)>2
        );