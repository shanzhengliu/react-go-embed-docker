import './App.css'
import useSWR from "swr";
import axios from "axios";
import { mockInit } from './mock/lib';
function App() {
  const fetcher = (url: string) => axios.get(url).then(res => res.data)
  if(import.meta.env.MODE === 'mock'){
    mockInit()
  }
  const { data } = useSWR(import.meta.env.VITE_BACKEND_URL, fetcher)
  return (
    <>
      <div>
          <h1 className="text-6xl">
              {data && data}
          </h1>
      </div>
    </>
  )
}

export default App
