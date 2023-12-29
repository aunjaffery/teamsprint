const Domain = import.meta.env.VITE_ENDPOINT
  ? import.meta.env.VITE_ENDPOINT
  : "http://localhost:8088";

export default Domain;
