/* Write your PL/SQL query statement below */

-- db type: Oracle
-- Level: Hard

select Department,Employee,salary from (
select t2.name Department,t.name Employee,t.salary,
dense_rank() OVER (PARTITION BY t.departmentid ORDER BY t.salary desc) pxh
	from Employee t,Department t2 where t.departmentid=t2.id
) where pxh <=3