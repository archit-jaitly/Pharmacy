# Pharmacy Inventory Management CRUD Assignment



This project utilizes Golang with the GoFiber framework and MongoDB to create an API for pharmacy inventory management. It allows users to add new medicines to the database, modify medicine details, and adjust stock quantities. Additionally, it incorporates an SMS controller to simulate SMS funcionality about newly added medicines to the pharmacy owner.
## Database Schema
```
type Medicine struct {
	ID           string `json:"id" bson:"_id,omitempty"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_no"`
	MoleculeName string `json:"molecule_name"`
	Quantity     int    `json:"quantity"`
}
```

## API Doc


### 1. Add Medicine details


- **HTTP Method** : POST  
- **URI** : ```/pharmacy/addMedicine```
- **REQUEST BODY** :
```
{
    "name": "Crocin",
    "molecule_name": "H2C5Na",
    "quantity": 86

}
``` 
### 2. Update Medicine details


- **HTTP Method** : PUT  
- **URI** : ```/pharmacy/updateMedicine/{id}```
- **REQUEST BODY** :
```
{
    "name": "Paracetamol",
    "molecule_name": "H2C5Na",
    "quantity": 4

}
``` 

### 3. Increment Medicine Stock


- **HTTP Method** : PUT  
- **URI** : ```/pharmacy/increment/{id}```
- **REQUEST BODY** : *Null*

### 4. Decrement Medicine Stock


- **HTTP Method** : PUT  
- **URI** : ```/pharmacy/decrement/{id}```
- **REQUEST BODY** : *Null*
 