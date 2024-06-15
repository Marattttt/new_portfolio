import { BrowserRouter, Route, Routes } from "react-router-dom"
import Editor from "./pages/Editor"

const AppRoutes = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Editor />} />
      </Routes>
    </BrowserRouter>
  )
}

export default AppRoutes
