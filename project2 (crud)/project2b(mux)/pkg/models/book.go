package models

import(
	"github.com/jinzhu/gorm"
	"github.com/a1nn1997/go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct{  //models for db tables
	gorm.Model
	Name string `gorm:""json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

type Library struct{  //models for db tables
	gorm.Model
	Name string `gorm:""json:"name"`
	Book string `json:"book"`
	Capacity string `json:"capacity"`
}

type Bookiya struct{  //models for db tables
		Name string `json:"name"`
		Author string `json:"author"`
		Publication string `json:"publication"`
	}

type Librarya struct{  //models for db tables
	Name string `json:"name"`
	Book *Bookiya `json:"book"`
	Capacity string `json:"capacity"`
}


func init(){    //db connect then get and than automigrate
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
	db.AutoMigrate(&Library{})
}

func (b *Book) CreateBook() *Book{
	db.NewRecord(b)     //first new record was taken then it was created
	db.Create(&b)
	return b
}

func GetAllBook() []Book{
	var Books []Book   //find all records as slices
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB){
	var getBook Book
	db:=db.Where("ID=?", Id).Find(&getBook)   //ID
	return &getBook, db
}

func DeleteBook(ID int64) Book{
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}

func (b *Library) CreateLibrary() *Library{
	db.NewRecord(b)     //first new record was taken then it was created
	db.Create(&b)
	return b
}

func GetAllLibrary() []Librarya{
	var Librarys []Library   //find all records as slices
	db.Find(&Librarys)
	var libraryas []Librarya
	for i,_ := range Librarys{
		var librarya Librarya
		librarya.Name=Librarys[i].Name
		librarya.Book=&Bookiya{Name:"", Author:"",Publication: ""}
		var Books []Book   //find all records as slices
		db.Find(&Books)
		for j,_ := range Books{
			if(Books[j].Name==Librarys[i].Book){
				librarya.Book.Name=Books[j].Name
				librarya.Book.Author=Books[j].Author
				librarya.Book.Publication=Books[j].Publication
			break;
			}
		}
		librarya.Capacity=Librarys[i].Capacity
		libraryas = append(libraryas, librarya)
	}
		return libraryas
}

func GetLibraryById(Id int64) (*Library, *gorm.DB){
	var getLibrary Library
	db:=db.Where("ID=?", Id).Find(&getLibrary)   //ID
	return &getLibrary, db
}

func DeleteLibrary(ID int64) Library{
	var Library Library
	db.Where("ID=?", ID).Delete(Library)
	return Library
}


func GetAllBooks() []Bookiya{
	var Books []Book   //find all records as slices
	db.Find(&Books)
	
	var Bookyas []Bookiya
	for i,_ := range Books{
		var Bookya Bookiya
		Bookya.Name =Books[i].Name
		Bookya.Author =Books[i].Author
		Bookya.Publication =Books[i].Publication
		Bookyas = append(Bookyas, Bookya)
	}
	return Bookyas
}

//Where Find Delete NewRecord Create