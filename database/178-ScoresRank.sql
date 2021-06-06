/* Write your PL/SQL query statement below */

-- db type: Oracle
-- Level: Medium

select score,
dense_rank() over(order by score desc) rank
from scores;