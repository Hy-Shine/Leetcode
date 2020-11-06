-- 184. Department Highest Salary
/* Write your PL/SQL query statement below */
select t2.name department,t.name employee,t.salary
from employee t,department t2 
where t.departmentid=t2.id 
and (t.departmentid,t.salary) in (select departmentid,max(salary) from employee group by departmentid)
order by t2.name;