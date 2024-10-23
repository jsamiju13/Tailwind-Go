def validarOpcion(opciones, mensaje): #INTRODUCE LA CANTIDAD DE OPCIONES QUE HAY, LUEGO EL MENSAJE QUE QUIERES QUE SE REPRODUZCA
    lista = []
    for i in range(opciones):
        lista.append(i + 1)
    while True:
        try:
            dato = int(input(mensaje))
            if dato in lista:
                return dato
            else:
                print("opcion incorrecta")
        except:
            print("valor incorrecto")


def validarNumero(mensaje): #INTRODUCE EL MENSAJE QUE QUIERAS QUE SE REPRODUZCA
    while True:
        try:
            numero = float(input(mensaje))
            return numero
        except ValueError:
            print("Eso no es un numero")


def validarRango(mensaje, iz=0, der=None): #INTRODUCE EL RANGO IZQUIERDO, LUEGO EL DERECHO Y LUEGO EL MENSAJE QUE QUIERAS QUE SE REPRODUZCA
    while True:
        try:
            numero = float(input(mensaje))
            if der == None:
                if numero < iz:
                    print("el nuermo que ha ingresado no puede ser menor a ", iz)
                else:
                    return numero

            elif iz <= numero <= der:
                return numero
            else:
                print("El Numero está fuera de rango")
        except ValueError:
            print("Eso no es un numero")

def validarRangoConocido(numero, iz=0, der=None): #INTRODUCE EL RANGO IZQUIERDO, LUEGO EL DERECHO Y LUEGO EL MENSAJE QUE QUIERAS QUE SE REPRODUZCA
    while True:
        try:
            if der == None:
                if numero < iz:
                    print("el nuermo que ha ingresado no puede ser menor a ", iz)
                else:
                    return numero

            elif iz <= numero <= der:
                return numero
            else:
                print("El Numero está fuera de rango")
                numero = int(input("Reingrese su numero\n"))
        except ValueError:
            print("Eso no es un numero")


def validarRangoInverso(min, max, mensaje): #INTRODUCE EL RANGO MINIMO, LUEGO EL MAXIMO Y LUEGO EL MENSAJE QUE QUIERAS QUE SE REPRODUZCA
    while True:
        try:
            numero = float(input(mensaje))
            if numero <= min or numero >= max:
                return numero
            else:
                print("ese numero no está en el rango")
        except ValueError:
            print("Eso no es un numero")

def validarDigito(mensaje):
    while True:
        cadena = input(mensaje)
        if cadena.isdigit():
            return cadena
        else:
            print("Su cadena solo debe contener digitos")

def validarLetra(mensaje):
    while True:
        nombre = input(mensaje)
        nom = nombre.split(" ")
        for i in nom:
            if not i.isalpha():
                print("El nombre ingresado contiene caracteres no permitidos")
                break
        else:
            return nombre
            break


def superficie(dim1=None,dim2=None,dim3=None):
    if dim2 == None and dim3 == None:
        return 3.14*dim1**2
    elif dim3 == None:
        return dim1*dim2
    else:
        return ((dim1+dim3)/2)*dim2
