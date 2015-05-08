package main

import (
	"log"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/fjl/go-couchdb"
)

func resourceCouchDB() *schema.Resource {
	return &schema.Resource{
		Create: resourceCouchCreate,
		Read:   resourceCouchRead,
		Update: resourceCouchUpdate,
		Delete: resourceCouchDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceCouchCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*couchdb.Client)

	//Create the database
	rec, err := client.CreateDB(d.Get("name").(string))

	if err != nil {
		return fmt.Errorf("Failed to create database: %s", err)
	}

	d.SetId(rec.Name())

	log.Printf("[INFO] database ID: %s", d.Id())

	return nil
}

func resourceCouchRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceCouchUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceCouchDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*couchdb.Client)

	log.Printf("[INFO] Deleting database: %s, %s", d.Get("name").(string), d.Id())

	err := client.DeleteDB(d.Get("name").(string))

	if err != nil {
		return fmt.Errorf("Error deleting database: %s", err)
	}

	return nil
}
