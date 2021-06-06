/* Write your PL/SQL query statement below */

-- db type: Oracle
-- Level: Medium

SELECT DISTINCT t.num ConsecutiveNums 
	FROM logs t,logs t2,logs t3 
WHERE t.id=(t2.id-1) AND t.num=t2.num 
	AND t2.id=(t3.id-1) AND t2.num=t3.num;

SELECT DISTINCT t.num ConsecutiveNums 
	FROM logs t 
	LEFT JOIN logs t2 ON t.id=(t2.id-1) AND t.NUM=t2.num
	LEFT JOIN logs t3 ON t2.id=(t3.id-1) AND t2.num=t3.num
WHERE t3.id IS NOT NULL;

-- 可以获得指定连续次数的数字
SELECT DISTINCT num ConsecutiveNums 
FROM ( 
	SELECT a.num,a.id-a.order1 FROM (
		SELECT t.*,ROW_NUMBER() OVER(PARTITION BY num ORDER BY id) order1
			FROM logs t) a GROUP BY num,(a.id-a.order1) HAVING COUNT(1)>2 
);