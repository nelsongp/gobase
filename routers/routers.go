package routers

import (
	"encoding/json"
	"github.com/nelsongp/gobase/bd"
	"github.com/nelsongp/gobase/models"
	"net/http"
)

func Registro(w http.ResponseWriter, r *http.Request){

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos " + err.Error(),400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Error contrasenia debe ser mayor a 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "ya existe un usuario con ese email", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al ingresar el registro de usuario " + err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro  del usuarrio", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
