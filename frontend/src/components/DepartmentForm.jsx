import { useState } from "react";

const DepartmentForm = () => {
    const [formData, setFormData] = useState({
        name: "",
        description: "",
        active: true,
    });
    const [message, setMessage] = useState(null);

    const handleChange = (e) => {
        const { name, value } = e.target;
        setFormData({ ...formData, [name]: value });
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            const response = await fetch("http://172.17.0.1:8080/department", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(formData),
            });

            if (!response.ok) throw new Error("Erro ao cadastrar departamento");

            setMessage({ type: "success", text: "Departamento cadastrado com sucesso!" });
        } catch (error) {
            setMessage({ type: "error", text: error.message });
        }
    };

    return (
        <div className="container mt-4">
            <h2>Cadastro de Departamento</h2>
            {message && (
                <div className={`alert ${message.type === "success" ? "alert-success" : "alert-danger"}`} role="alert">
                    {message.text}
                </div>
            )}
            <form onSubmit={handleSubmit}>
                <div className="mb-3">
                    <label className="form-label">Nome</label>
                    <input
                        type="text"
                        className="form-control"
                        name="name"
                        value={formData.name}
                        onChange={handleChange}
                        required
                    />
                </div>
                <div className="mb-3">
                    <label className="form-label">Descrição</label>
                    <textarea
                        className="form-control"
                        name="description"
                        value={formData.description}
                        onChange={handleChange}
                        required
                    />
                </div>
                <div className="mb-3 form-check">
                    <input
                        type="checkbox"
                        className="form-check-input"
                        name="active"
                        checked={formData.active}
                        onChange={() => setFormData({ ...formData, active: !formData.active })}
                    />
                    <label className="form-check-label">Ativo</label>
                </div>
                <button type="submit" className="btn btn-primary">Cadastrar</button>
            </form>
        </div>
    );
};

export default DepartmentForm;
