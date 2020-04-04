
# Desplegar Cluster de OpenShift desde Terraform

_Esta guia es un paso a  paso del despliegue de un Cluster de Openshift version3.11 en IBM Cloud desde CLI, utilizando terraform._


## Comenzando 🚀

_Para iniciar con el despliegue es necesario que hagamos un pull de la imagen Docker provista por ibm en la cual se encuantran las diferentes plantillas y configuracion para el despliegue del cluster._

_El pull de esta imagen docker se puede hacer utilizando el siguiente comando:_

    
    docker pull ibmterraform/terraform-provider-ibm-docker
    
    
_Despues de haber hecho el pull de la imagen, debemos verificar que se a agregado correctamente la imagen y esta verificación se hace mediente el siguinete comando:_

    
    docker images
    

_Ahora debemos ingresar en esta imagen que acabamos de subir para desde ahi realizar el despliegue, Para ingresar a la imagen debemos ingresar el siguiente comando:_

    
    docker run -it ibmterraform/terraform-provider-ibm-docker:latest /bin/bash
    
    
## Despliegue 📦

_Apenas ingresemos a la imagen debe aparecer al inicio de la linea de comandos lo siguiente:_

    Bash-4.4#
    
_Ahora debemos clonar el repositorio de GitHub, y esto se hace mediante el siguiente comando:_

    git clone https://github.com/emeloibmco/Schematics-Classic-Infrastructure-ROKS.git
    
_Al tener clonado el repositorio, debemos ingresar a este preyecto y lo hacemos mediante el siguiente comando:_

    cd terraform-openshift/ibm-iks-openshift
 
## Ahora debemos modificar los diferentes archivos que encontramos aqui 🛠️

### main.tf

_Para poder entrar a editar este archivo debemos ingresar el siguinete comando:_

    vi main.tf
    
_Apenas ingresemos al archivo debemos precionar la tecla **i** del teclado, al hacer esto se nos habilitara la edición del archivo._

_Los campos que debemos modificar de este archivos son:_

    default_pool_size =  1
    resource_group_id =  " xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx "
    
_**Notas:**_ 

_1. El campo **defauld_pool_zise** hace referencia al numero de nodos con el cual se va a crear el cluster._

_2. El campo **resource_group_id** hace referencia al ID del Grupo de recursos en el cual queremos crear el cluster._

_2.1 Para conocer el ID del los diferentes grupos de recursos que tenemos en IBM Cloud debemos colocar en **otra terminal** el siguiente comando:_

    ibmcloud resource groups
    
_Para salir y guardar los cambios hechos en este documento de debe precionar primero la tecla **esc** y luego digitar **:wq!** y por ultimo dar enter._ 

### provider.tf

_Para poder entrar a editar este archivo debemos ingresar el siguinete comando:_

    vi provider.tf
    
_Apenas ingresemos al archivo debemos precionar la tecla **i** del teclado, al hacer esto se nos habilitara la edición del archivo._

_En este archivo solo se debe modificar un campo el cual es:_

    region =  "us-south"
    
_**Nota:**_ 

_1. El campo **region** hace referencia al region en la cual desea crear el cluster._

    
_Para salir y guardar los cambios hechos en este documento de debe precionar primero la tecla **esc** y luego digitar **:wq!** y por ultimo dar enter._ 

### terraform.tfvars

_Para poder entrar a editar este archivo debemos ingresar el siguinete comando:_

    vi terraform.tfvars
    
_Apenas ingresemos al archivo debemos precionar la tecla **i** del teclado, al hacer esto se nos habilitara la edición del archivo._

_Los campos que debemos modificar de este archivos son:_

    iaas_classic_api_key =  "xxxxxxxxxxxxxxxxxxxxxxxxxx"
    iaas_classic_username =  "xxxxxxxx_usuario@ibm.com"
    ibmcloud_api_key =  "xxxxxxxxxxxxxxxxxxxxxxxx"
    
_**Notas:**_ 

_1. El campo **iaas_classic_api_key** hace referencia al api-key de infraestructura clasica que se pedia como prerequisito._

_2. El campo **iaas_classic_username** hace referencia al username que se compone asi: <numerocuenta_correo>._

_3. El campo **ibmcloud_api_key** hace referencia al api-key de ibmcloud que se pedia como prerequisito._
    
_Para salir y guardar los cambios hechos en este documento de debe precionar primero la tecla **esc** y luego digitar **:wq!** y por ultimo dar enter._ 

### variables.tf

_Para poder entrar a editar este archivo debemos ingresar el siguinete comando:_

    vi variables.tf
    
_Apenas ingresemos al archivo debemos precionar la tecla **i** del teclado, al hacer esto se nos habilitara la edición del archivo._

_Los campos que debemos modificar de este archivos son:_

    variable "datacenter"
    variable "private_vlan_id"
    variable "public_vlan_id"
        
_**Notas:**_ 

_1. El campo **datacenter** hace referencia a la zona donde esta ubicado el datacenter en el cual queremos que este alojado nuestro cluster, esta zona debe coincidir con la region que seleccionamos en el archivo provider.tf._

_2. El campo **private_vlan_id** hace referencia a la vlan privada que se tenia que crear como prerequisito._

_3. El campo **public_vlan_id** hace referencia a la vlan publica que se tenia que crear como prerequisito._

    
_Para salir y guardar los cambios hechos en este documento de debe precionar primero la tecla **esc** y luego digitar **:wq!** y por ultimo dar enter._ 

## Despliegue del CLuster con terraform 🖇️

_Despues de haber modificado los archivos anteriores, debemos digitar el siguiente comando:_

    terraform init

_Apenas termine de ejecutarse el comando anterior, debemos digitar el siguiente comando:_

    terraform plan
    
_En el comando anterior se va a demorar un poco ya que esta contruyendo el plano para el despliegue. Al termiar de ejecutar este comando se debe digitar el siguiente:_

    terraform apply
    
_Al ejecutar el comando anterior, despues de unos segundos nos va a parecer una pregunta en la cual debemos contestar **yes**, y aca estamos dando la autorización de crear el cluster de OpenShift en IBM Cloud._



### Pre-requisitos 📋

_Para el correcto funcionameniento de este despliegue en IBM Cloud son necesarios los siguinetes programas:_

* _Docker:_

    ```
    sudo apt install docker-ce
    ```

* _Ansible:_
    ```
    pip install "ansible>=2.8.0"
    ```
_Ademas debemos crear desde la pagina de IBM Cloud los siguientes recursos:_

* 1 Api_key de IBM Cloud.
* 1 Api_key para infraestructura clasica.
* 1 Vlan publica.
* 1 Vlan privada.


## Autor ✒️

* **Jhoiner Smith Rojas González** -  - jhoinerrojas@ibm.com
_Equipo IBM Cloud Tech Sales Colombia._


