import { useForm } from "react-hook-form";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";
import { Eye, EyeOff } from "lucide-react"; // Import ikon mata

export default function LoginPage() {
  const {
    register,
    handleSubmit,
    reset,
    formState: { errors },
  } = useForm();
  const [loading, setLoading] = useState(false);
  const [errorMessage, setErrorMessage] = useState("");
  const [typePage, setTypePage] = useState("login");
  const [showPassword, setShowPassword] = useState(false);
  const navigate = useNavigate();
  const apiUrl = import.meta.env.VITE_API_URL;

  const changeType = () => {
    setTypePage(typePage === "login" ? "register" : "login");
  };

  const onSubmit = async (data) => {
    setLoading(true);
    setErrorMessage("");

    try {
      if (typePage === "login") {
        const response = await axios.post(`${apiUrl}login`, data);
        localStorage.setItem("token", response.data.user);
        navigate("/");
      } else {
        const response = await axios.post(`${apiUrl}register`, data);
        reset();
        setTypePage("login");
      }
    } catch (error) {
      setErrorMessage("Invalid username or password");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <div className="bg-white p-8 shadow-md rounded-lg w-96">
        <h2 className="text-2xl font-semibold mb-4 text-center">
          {typePage === "login" ? "Login" : "Register"}
        </h2>
        {errorMessage && (
          <p className="text-red-500 text-center">{errorMessage}</p>
        )}

        <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
          <div>
            <label className="block text-gray-700">Username</label>
            <input
              type="text"
              {...register("username", { required: "Username is required" })}
              className="w-full px-3 py-2 border rounded"
            />
            {errors.username && (
              <p className="text-red-500 text-sm">{errors.username.message}</p>
            )}
          </div>
          <div className="relative">
            <label className="block text-gray-700">Password</label>
            <div className="relative">
              <input
                type={showPassword ? "text" : "password"}
                {...register("password", { required: "Password is required" })}
                className="w-full px-3 py-2 border rounded pr-10"
              />
              <button
                type="button"
                onClick={() => setShowPassword(!showPassword)}
                className="absolute inset-y-0 right-0 flex items-center pr-3"
              >
                {showPassword ? (
                  <EyeOff className="h-5 w-5 text-gray-500" />
                ) : (
                  <Eye className="h-5 w-5 text-gray-500" />
                )}
              </button>
            </div>
            {errors.password && (
              <p className="text-red-500 text-sm">{errors.password.message}</p>
            )}
          </div>
          <button
            type="submit"
            disabled={loading}
            className="w-full bg-blue-600 text-white py-2 rounded hover:bg-blue-700 disabled:bg-gray-400"
          >
            {loading
              ? typePage === "login"
                ? "Logging in..."
                : "Registering..."
              : typePage === "login"
              ? "Login"
              : "Register"}
          </button>
        </form>
        <p className="mt-3 cursor-pointer" onClick={changeType}>
          {typePage === "login"
            ? "Belum Punya Akun ?, Register"
            : "Sudah Punya Akun ?, Login"}
        </p>
      </div>
    </div>
  );
}
