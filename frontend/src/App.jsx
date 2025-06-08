import { useState } from "react";
import EmployeeForm from "./components/EmployeeForm";
import EmployeeList from "./components/EmployeeList";
import DepartmentForm from "./components/DepartmentForm";
import "bootstrap/dist/css/bootstrap.min.css";
import "./App.css";

function App() {
  const [activeView, setActiveView] = useState("employee-form");

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
          onClick={() => setActiveView("employee-list")}
          style={menuStyle(activeView === "employee-list")}
        >
          Lista de Funcionários
        </button>

        <button
          onClick={() => setActiveView("employee-form")}
          style={menuStyle(activeView === "employee-form")}
        >
          Cadastrar Funcionário
        </button>

        <button
          onClick={() => setActiveView("department")}
          style={menuStyle(activeView === "department")}
        >
          Cadastrar Departamento
        </button>
      </aside>

      {/* Área de Conteúdo */}
      <main
        style={{
          flexGrow: 1,
          backgroundColor: "#f0f0f0",
          padding: "2rem",
          overflowX: "auto",
          width: "1600px",
        }}
      >
        {activeView === "employee-form" && <EmployeeForm />}
        {activeView === "employee-list" && <EmployeeList />}
        {activeView === "department" && <DepartmentForm />}
      </main>
    </div>
  );
}

const menuStyle = (isActive) => ({
  backgroundColor: isActive ? "#808000" : "#333",
  color: "#fff",
  border: "none",
  width: "100%",
  padding: "1rem",
  marginBottom: "1rem",
  cursor: "pointer"
});

export default App;
