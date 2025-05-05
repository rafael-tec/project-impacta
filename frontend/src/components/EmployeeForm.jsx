import { useState } from "react";

const EmployeeForm = () => {
  const [formData, setFormData] = useState({
    name: "",
    age: "",
    salary: "",
    hiring_date: "",
    dismissal_date: "",
    department_id: "",
    job_title: "",
    active: true
  });

  const [message, setMessage] = useState(null);

  const handleChange = (e) => {
    const { name, value, type, checked } = e.target;
    setFormData({
      ...formData,
      [name]: type === "checkbox" ? checked : value,
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const response = await fetch("http://172.17.0.1:8080/employee", {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify(formData)
      });

      if (response.status === 201) {
        setMessage({ type: "success", text: "Funcionário cadastrado com sucesso!" });
      } else {
        setMessage({ type: "error", text: "Erro ao cadastrar funcionário." });
      }
    } catch (error) {
      setMessage({ type: "error", text: "Erro na requisição: " + error.message });
    }
  };

  return (
    <div className="container mt-4">
      <h2>Cadastro de Funcionário</h2>
      {message && (
        <div className={`alert ${message.type === "success" ? "alert-success" : "alert-danger"}`} role="alert">
          {message.text}
        </div>
      )}
      <form onSubmit={handleSubmit}>
        <div className="mb-3">
          <label className="form-label">Nome</label>
          <input type="text" className="form-control" name="name" value={formData.name} onChange={handleChange} required />
        </div>
        <div className="mb-3">
          <label className="form-label">Idade</label>
          <input type="number" className="form-control" name="age" value={formData.age} onChange={handleChange} required />
        </div>
        <div className="mb-3">
          <label className="form-label">Salário</label>
          <input type="number" step="0.01" className="form-control" name="salary" value={formData.salary} onChange={handleChange} required />
        </div>
        <div className="mb-3">
          <label className="form-label">Data de Admissão</label>
          <input type="date" className="form-control" name="hiring_date" value={formData.hiring_date} onChange={handleChange} required />
        </div>
        <div className="mb-3">
          <label className="form-label">ID do Departamento</label>
          <input type="number" className="form-control" name="department_id" value={formData.department_id} onChange={handleChange} required />
        </div>
        <div className="mb-3">
          <label className="form-label">Cargo</label>
          <input type="text" className="form-control" name="job_title" value={formData.job_title} onChange={handleChange} required />
        </div>
        <div className="mb-3 form-check">
          <input type="checkbox" className="form-check-input" name="active" checked={formData.active} onChange={handleChange} />
          <label className="form-check-label">Ativo</label>
        </div>
        <button type="submit" className="btn btn-primary">Cadastrar Funcionário</button>
      </form>
    </div>
  );
};

export default EmployeeForm;
