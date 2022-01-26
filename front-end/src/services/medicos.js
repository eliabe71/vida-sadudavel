import { api } from './config'

export default {
  listar:()=>{
    return api.get('medics')
  },
  getMedico:(id)=>{
    return api.get(`getmedic/${id}`)
  }
}