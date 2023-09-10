import './App.css'
import useSWR from "swr";
import axios from "axios";
function App() {
  const fetcher = (url: string) => axios.get(url).then(res => res.data)
  const { data } = useSWR(import.meta.env.VITE_BACKEND_URL, fetcher)
  return (
    <>
      <div>
          <h1 className="text-6xl">
              {data}
          </h1>
      </div>
    </>
  )
}

export default App
