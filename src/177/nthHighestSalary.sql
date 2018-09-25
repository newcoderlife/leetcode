CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT BEGIN declare P INT;
set
  P = N - 1;
RETURN (
    select
      (
        select
          DISTINCT Salary
        from
          Employee
        order by
          Salary DESC
        limit
          P, 1
      )
  );
END