import { useEffect, useState } from "react";
import { useForm } from "react-hook-form";
import { useNavigate, useParams } from "react-router-dom";
import axios from "axios";

export default function AddEdit() {
  const { id } = useParams();
  const navigate = useNavigate();
  const isEdit = id !== "new";
  const apiUrl = import.meta.env.VITE_API_URL;

  const {
    register,
    handleSubmit,
    setValue,
    formState: { errors, isSubmitting },
  } = useForm();

  useEffect(() => {
    if (isEdit) {
      axios
        .get(`${apiUrl}product/${id}`)
        .then((res) => {
          const { name, sku, quantity, location, status } = res.data.data;
          console.log(res.data);

          setValue("name", name);
          setValue("sku", sku);
          setValue("quantity", quantity);
          setValue("location", location);
          setValue("status", status);
        })
        .catch((err) => console.error("Error fetching product:", err));
    }
  }, [id, isEdit, setValue]);

  const onSubmit = async (data) => {
    try {
      if (isEdit) {
        await axios.put(`${apiUrl}product/${id}`, {
          name: data.name,
          sku: data.sku,
          quantity: parseInt(data.quantity),
          location: data.location,
          status: data.status,
        });
      } else {
        await axios.post(`${apiUrl}}product`, {
          name: data.name,
          sku: data.sku,
          quantity: parseInt(data.quantity),
          location: data.location,
          status: data.status,
        });
      }
      navigate("/");
    } catch (error) {
      console.error("Error saving product:", error);
    }
  };

  return (
    <div className="p-6 bg-gray-100 min-h-screen">
      <h2 className="text-2xl font-semibold mb-4">
        {isEdit ? "Edit Product" : "Add New Product"}
      </h2>
      <form
        onSubmit={handleSubmit(onSubmit)}
        className="bg-white p-6 rounded-lg shadow-md"
      >
        {/* Name */}
        <div className="mb-4">
          <label className="block text-gray-700">Name</label>
          <input
            type="text"
            {...register("name", { required: "Name is required" })}
            className="w-full border p-2 rounded"
          />
          {errors.Name && (
            <p className="text-red-500 text-sm">{errors.Name.message}</p>
          )}
        </div>

        {/* SKU */}
        <div className="mb-4">
          <label className="block text-gray-700">SKU</label>
          <input
            type="text"
            {...register("sku", { required: "SKU is required" })}
            className="w-full border p-2 rounded"
          />
          {errors.SKU && (
            <p className="text-red-500 text-sm">{errors.SKU.message}</p>
          )}
        </div>

        {/* Quantity */}
        <div className="mb-4">
          <label className="block text-gray-700">Quantity</label>
          <input
            type="number"
            {...register("quantity", {
              required: "Quantity is required",
              min: { value: 1, message: "Quantity must be at least 1" },
            })}
            className="w-full border p-2 rounded"
          />
          {errors.Quantity && (
            <p className="text-red-500 text-sm">{errors.Quantity.message}</p>
          )}
        </div>

        {/* Location */}
        <div className="mb-4">
          <label className="block text-gray-700">Location</label>
          <input
            type="text"
            {...register("location", { required: "Location is required" })}
            className="w-full border p-2 rounded"
          />
          {errors.Location && (
            <p className="text-red-500 text-sm">{errors.Location.message}</p>
          )}
        </div>

        {/* Status */}
        <div className="mb-4">
          <label className="block text-gray-700">Status</label>
          <select
            {...register("status", { required: "Status is required" })}
            className="w-full border p-2 rounded"
          >
            <option value="">Select Status</option>
            <option value="Available">Available</option>
            <option value="Out of Stock">Out of Stock</option>
          </select>
          {errors.Status && (
            <p className="text-red-500 text-sm">{errors.Status.message}</p>
          )}
        </div>

        {/* Buttons */}
        <div className="flex space-x-4">
          <button
            type="submit"
            disabled={isSubmitting}
            className="px-4 py-2 bg-blue-600 text-white rounded disabled:bg-blue-300"
          >
            {isSubmitting ? "Saving..." : isEdit ? "Update" : "Create"}
          </button>
          <button
            type="button"
            onClick={() => navigate("/")}
            className="px-4 py-2 bg-gray-400 text-white rounded"
          >
            Cancel
          </button>
        </div>
      </form>
    </div>
  );
}
