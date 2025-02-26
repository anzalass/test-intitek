import React from "react";
import ReactDOM from "react-dom/client";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import AddEdit from "./page/AddEdit.jsx";
import Homepage from "./page/Homepage.jsx";
import LoginPage from "./page/Loginpage.jsx";
import ProtectedRoute from "./utils/ProtectedRoute.jsx"; // Import ProtectedRoute
import AuthGuard from "./utils/AuthGuard.jsx";

ReactDOM.createRoot(document.getElementById("root")).render(
  <React.StrictMode>
    <Router>
      <Routes>
        {/* Halaman Login (Bisa diakses tanpa login) */}
        <Route element={<AuthGuard />}>
          <Route path="/auth" element={<LoginPage />} />
        </Route>
        {/* Halaman yang Butuh Login */}
        <Route element={<ProtectedRoute />}>
          <Route path="/" element={<Homepage />} />
          <Route path="/:id" element={<AddEdit />} />
        </Route>
      </Routes>
    </Router>
  </React.StrictMode>
);
