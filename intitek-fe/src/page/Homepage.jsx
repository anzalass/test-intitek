import { useEffect, useState } from "react";
import axios from "axios";
import { Link, useNavigate } from "react-router-dom";

export default function Homepage() {
  const [products, setProducts] = useState([]);
  const [statusFilter, setStatusFilter] = useState(""); // Filter Status
  const [lowStockFilter, setLowStockFilter] = useState(""); // Filter Low Stock
  const [currentPage, setCurrentPage] = useState(1); // Pagination
  const [pageSize, setPageSize] = useState(5); // Ukuran halaman (bisa diubah)
  const [totalPages, setTotalPages] = useState(1); // Total Halaman
  const navigate = useNavigate();
  const apiUrl = import.meta.env.VITE_API_URL;

  useEffect(() => {
    fetchData();
  }, [statusFilter, lowStockFilter, currentPage, pageSize]);

  const handleDelete = async (id) => {
    try {
      await axios.delete(`${apiUrl}product/${id}`);
      fetchData();
    } catch (error) {
      console.error("Error deleting product:", error);
    }
  };

  async function fetchData() {
    try {
      const res = await axios.get(`${apiUrl}products`, {
        params: {
          status: statusFilter,
          page: currentPage,
          pageSize: pageSize,
          low_stock: lowStockFilter,
        },
      });

      setProducts(res.data.data);
      setTotalPages(Math.ceil(res.data.total / pageSize)); // Hitung total halaman
    } catch (error) {
      console.error("Error fetching data:", error);
    }
  }

  const logout = () => {
    localStorage.removeItem("token");
    window.location.href = "/auth";
  };
  return (
    <div className="p-6 bg-gray-100 min-h-screen">
      {/* Filter Section */}
      <div className="mb-4   gap-6 flex justify-between w-full">
        <div className="flex  gap-6 flex-wrap">
          <div>
            <h3 className="font-medium mb-2">Filter by Status</h3>
            <select
              value={statusFilter}
              onChange={(e) => setStatusFilter(e.target.value)}
              className="border rounded px-3 py-1"
            >
              <option value="">All</option>
              <option value="Available">Available</option>
              <option value="Out of Stock">Out of Stock</option>
            </select>
          </div>

          {/* Filter Low Stock */}
          <div>
            <h3 className="font-medium mb-2">Filter by Low Stock</h3>
            <select
              value={lowStockFilter}
              onChange={(e) => setLowStockFilter(e.target.value)}
              className="border rounded px-3 py-1"
            >
              <option value="">All</option>
              <option value="true">Low Stock</option>
              <option value="false">Sufficient Stock</option>
            </select>
          </div>

          {/* Page Size */}
          <div>
            <h3 className="font-medium mb-2">Page Size</h3>
            <select
              value={pageSize}
              onChange={(e) => setPageSize(Number(e.target.value))}
              className="border rounded px-3 py-1"
            >
              {[5, 10, 15, 20].map((size) => (
                <option key={size} value={size}>
                  {size}
                </option>
              ))}
            </select>
          </div>
        </div>
        <div className="text-white space-x-2">
          <Link to={"/new"}>
            <button>Tambah Data +</button>
          </Link>

          <button onClick={logout}>Logout </button>
        </div>
        {/* Filter Status */}
      </div>

      {/* Table */}
      <div className="overflow-x-auto">
        <table className="min-w-full bg-white shadow-md rounded-lg overflow-hidden">
          <thead className="bg-gray-200 text-gray-600 uppercase text-sm">
            <tr>
              <th className="py-3 px-6 text-left">Name</th>
              <th className="py-3 px-6 text-left">SKU</th>
              <th className="py-3 px-6 text-left">Quantity</th>
              <th className="py-3 px-6 text-left">Location</th>
              <th className="py-3 px-6 text-left">Status</th>
              <th className="py-3 px-6 text-left">Actions</th>
            </tr>
          </thead>
          <tbody>
            {products.length > 0 ? (
              products.map((product) => (
                <tr
                  key={product.SKU}
                  className="border-b hover:bg-gray-100 transition"
                >
                  <td className="py-3 px-6">{product.name}</td>
                  <td className="py-3 px-6">{product.sku}</td>
                  <td className="py-3 px-6">{product.quantity}</td>
                  <td className="py-3 px-6">{product.location}</td>
                  <td className="py-3 px-6">
                    <span
                      className={`px-3 py-1 rounded-full text-xs font-semibold ${
                        product.status === "Available"
                          ? "bg-green-200 text-green-800"
                          : "bg-red-200 text-red-800"
                      }`}
                    >
                      {product.status}
                    </span>
                  </td>
                  <td className="py-3 px-6 flex space-x-2">
                    <button
                      onClick={() => navigate(`/${product.sku}`)}
                      className="px-3 py-1 bg-yellow-500 text-white rounded text-sm"
                    >
                      Edit
                    </button>
                    <button
                      onClick={() => handleDelete(product.sku)}
                      className="px-3 py-1 bg-red-600 text-white rounded text-sm"
                    >
                      Delete
                    </button>
                  </td>
                </tr>
              ))
            ) : (
              <tr>
                <td colSpan="6" className="text-center py-4 text-gray-500">
                  No products found.
                </td>
              </tr>
            )}
          </tbody>
        </table>
      </div>

      {/* Pagination */}
      <div className="flex justify-center mt-4 space-x-2">
        <button
          className={`px-4 py-2 text-white border rounded ${
            currentPage === 1
              ? "opacity-100 cursor-not-allowed"
              : "hover:bg-gray-200"
          }`}
          onClick={() => setCurrentPage((prev) => Math.max(prev - 1, 1))}
          disabled={currentPage === 1}
        >
          Prev
        </button>
        <span className="px-4 py-2 border rounded">
          Page {currentPage} of {totalPages}
        </span>
        <button
          className={`px-4 text-white py-2 border rounded ${
            currentPage >= totalPages
              ? "opacity-100 cursor-not-allowed"
              : "hover:bg-gray-200"
          }`}
          onClick={() => setCurrentPage((prev) => prev + 1)}
          disabled={currentPage >= totalPages}
        >
          Next
        </button>
      </div>
    </div>
  );
}
