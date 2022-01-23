import Login from '../views/Login.vue'
import PacienteHome from '../views/Paciente/PacienteHome.vue'
import PacienteMinhasConsultas from '../views/Paciente/PacienteMinhasConsultas.vue'
import PacienteEditarConsulta from '../views/Paciente/PacienteEditarConsulta.vue'

export default [
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/pacienteHome',
    name: 'PacienteHome',
    component: PacienteHome
  },
  {
    path: '/paciente/minhasConsultas',
    name: 'PacienteConsultas',
    component: PacienteMinhasConsultas
  },
  {
    path: 'paciente/minhasConsultas/edit/:id',
    name: 'pacienteEditarConsulta',
    component: PacienteEditarConsulta,
    params: true
  }
]