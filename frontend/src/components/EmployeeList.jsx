import { useEffect, useState } from "react";

const EmployeeList = () => {
  const [employees, setEmployees] = useState([]);
  const [filter, setFilter] = useState("");
  const [message, setMessage] = useState(null);

  useEffect(() => {
    const fetchEmployees = async () => {
      try {
        const response = await fetch("http://localhost:8080/employees");
        if (!response.ok) throw new Error("Erro ao carregar funcionários.");
        const data = await response.json();
        setEmployees(data);
      } catch (error) {
        setMessage({ type: "error", text: error.message });
      }
    };

    fetchEmployees();
  }, []);

  const dismissEmployee = async (id) => {
    const dismissalDate = new Date().toISOString().split("T")[0];
    try {
      const response = await fetch(`http://localhost:8080/employee?dismissal_date=${dismissalDate}&employee_id=${id}`, {
        method: "PATCH"
      });

      if (!response.ok) throw new Error("Erro ao demitir funcionário.");

      const updated = employees.map(emp =>
        emp.id === id ? { ...emp, active: false, dismissal_date: dismissalDate } : emp
      );
      setEmployees(updated);
      setMessage({ type: "success", text: "Funcionário demitido com sucesso." });
    } catch (err) {
      setMessage({ type: "error", text: err.message });
    }
  };

  const formatDate = (isoDate) => {
    if (!isoDate) return "-";
    const [year, month, day] = isoDate.split("-");
    return `${day}/${month}/${year}`;
  };

  const filtered = employees.filter(emp =>
    emp.name.toLowerCase().includes(filter.toLowerCase()) ||
    emp.job_title.toLowerCase().includes(filter.toLowerCase())
  );

  return (
    <div className="d-flex justify-content-center align-items-center" style={{ padding: "2rem" }}>
      <div style={{ width: "100%", maxWidth: "1200px" }}>
        <h2 className="text-center mb-4">Funcionários</h2>

        {message && (
          <div className={`alert ${message.type === "success" ? "alert-success" : "alert-danger"}`} role="alert">
            {message.text}
          </div>
        )}

        <input
          type="text"
          className="form-control mb-3"
          placeholder="Filtrar por nome ou cargo"
          value={filter}
          onChange={(e) => setFilter(e.target.value)}
        />

        <div className="table-responsive">
          <table className="table table-striped w-100">
            <thead>
              <tr>
                <th>Nome</th>
                <th>Cargo</th>
                <th>Status</th>
                <th>Data de Admissão</th>
                <th>Data de Demissão</th>
                <th>Ações</th>
              </tr>
            </thead>
            <tbody>
              {filtered.map(emp => (
                <tr key={emp.id} className={!emp.active ? "text-muted" : ""}>
                  <td>{emp.name}</td>
                  <td>{emp.job_title}</td>
                  <td>{emp.active ? "Ativo" : "Demitido"}</td>
                  <td>{formatDate(emp.hiring_date)}</td>
                  <td>{formatDate(emp.dismissal_date)}</td>
                  <td>
                    {emp.active && (
                      <button className="btn btn-danger btn-sm" onClick={() => dismissEmployee(emp.id)}>
                        Demitir
                      </button>
                    )}
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
};

export default EmployeeList;
