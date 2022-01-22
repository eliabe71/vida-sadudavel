import Login from '../views/Login.vue'
import PacienteHome from '../views/Paciente/PacienteHome.vue'

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
  }
]