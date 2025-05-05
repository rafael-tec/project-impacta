import { useState } from "react";
import DepartmentForm from "./components/DepartmentForm";
import EmployeeForm from "./components/EmployeeForm";
import "bootstrap/dist/css/bootstrap.min.css";
import "./App.css";

function App() {
  const [activeView, setActiveView] = useState("employee");

  return (
    <div className="app-layout" style={{ display: "flex", height: "100vh", fontFamily: "Arial, sans-serif" }}>
      {/* Menu Lateral */}
      <aside style={{
        width: "250px",
        backgroundColor: "#000",
        color: "#fff",
        padding: "2rem 1rem",
        display: "flex",
        flexDirection: "column",
        alignItems: "center"
      }}>
        {/* Logo/Ícone */}
        <div style={{
          width: 80,
          height: 80,
          backgroundColor: "#808000",
          borderRadius: "50%",
          marginBottom: "2rem",
          display: "flex",
          alignItems: "center",
          justifyContent: "center",
          fontSize: "2rem",
          fontWeight: "bold"
        }}>
          RD
        </div>

        <button
          onClick={() => setActiveView("employee")}
          style={{
            backgroundColor: activeView === "employee" ? "#808000" : "#333",
            color: "#fff",
            border: "none",
            width: "100%",
            padding: "1rem",
            marginBottom: "1rem",
            cursor: "pointer"
          }}
        >
          Cadastrar Funcionário
        </button>
        <button
          onClick={() => setActiveView("department")}
          style={{
            backgroundColor: activeView === "department" ? "#808000" : "#333",
            color: "#fff",
            border: "none",
            width: "100%",
            padding: "1rem",
            cursor: "pointer"
          }}
        >
          Cadastrar Departamento
        </button>
      </aside>

      {/* Área de Conteúdo */}
      <main style={{ flexGrow: 1, backgroundColor: "#f0f0f0", padding: "2rem" }}>
        {activeView === "employee" && <EmployeeForm />}
        {activeView === "department" && <DepartmentForm />}
      </main>
    </div>
  );
}

export default App;
