package controllers 

type ControllerInterface interface {
    Init() //Initialize the context and subclass name
    Prepare()                    //some processing before execution begins
    Get()                        //method = GET processing
    Post()                       //method = POST processing
    Delete()                     //method = DELETE processing
    Put()                        //method = PUT handling
    Head()                       //method = HEAD processing
    Patch()                      //method = PATCH treatment
    Options()                    //method = OPTIONS processing
    Finish()                     //executed after completion of treatment
    Render() error               //method executed after the corresponding method to render the page
}