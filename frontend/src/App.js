import logo from './logo.svg';
import { BrowserRouter, Routes, Route } from 'react-router-dom'

// Paginas
import Login from './pages/login';
import Cadastro from './pages/cadastro';
import Perfil from './pages/perfil';

import PaginaInicial from './pages/paginaInicial';
import Partida from './pages/partida';
import Resultados from './pages/resultados';

import Config from './pages/config'
import Header from './components/header';

import { AuthProvider } from './components/auth';
import Jogo from './pages/jogo';
import Ws from './pages/ws';

function App() {
  return (
    <BrowserRouter>
      <AuthProvider>

        <Header />
        <Routes>
          <Route path="/"
            element={<PaginaInicial />} />

          <Route path="/login"
            element={<Login />} />
          <Route path="/cadastro"
            element={<Cadastro />} />
          <Route path="/perfil"
            element={<Perfil />} />

          <Route path="/partida"
            element={<Partida />} />
          <Route path="/resultados"
            element={<Resultados />} />
          <Route path="/config"
            element={<Config />} />

          <Route path="/jogo"
            element={<Jogo />} />

          <Route path="/ws"
            element={<Ws />} />

          {/*Exemplo para privar as rotas no futuro*/}
          {/* <Route path="/admin/usuarios" */}
          {/*   element={<PrivateRoute allowedRoles={["admin"]}> <AdminUsuarios /> </PrivateRoute>} /> */}
        </Routes>
      </AuthProvider>
    </BrowserRouter>
  );
}

export default App;
