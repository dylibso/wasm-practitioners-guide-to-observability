#![allow(dead_code)]
use observe_api::*;
use serde::Deserialize;

const PRODUCT_JSON: &str = r#"{"id":1,"title":"iPhone 9","description":"An apple mobile which is nothing like apple","price":549,"discountPercentage":12.96,"rating":4.69,"stock":94,"brand":"Apple","category":"smartphones","thumbnail":"https://cdn.dummyjson.com/product-images/1/thumbnail.jpg","images":["https://cdn.dummyjson.com/product-images/1/1.jpg","https://cdn.dummyjson.com/product-images/1/2.jpg","https://cdn.dummyjson.com/product-images/1/3.jpg","https://cdn.dummyjson.com/product-images/1/4.jpg","https://cdn.dummyjson.com/product-images/1/thumbnail.jpg"]}"#;

#[derive(Debug, Default)]
struct Data {
    pub name: String,
    pub product_item: Product,
}

#[derive(Debug, Default, Deserialize)]
struct Product {
    pub id: usize,
    pub title: String,
    pub description: String,
    pub price: usize,
    #[serde(rename = "discountPercentage")]
    pub discount_percentage: f64,
    pub rating: f64,
    pub stock: usize,
    pub brand: String,
    pub category: String,
    pub thumbnail: String,
    pub images: Vec<String>,
}

impl Data {
    #[no_mangle]
    fn set_name(&mut self, name: &str) {
        self.name = String::from(name);
    }

    #[no_mangle]
    fn set_product(&mut self, input: &str) {
        let product: Product = serde_json::from_str(input).unwrap_or_default();

        if product.brand != "Apple" {
            span_tags(vec!["brand:unknown"])
        }

        span_tags(vec![
            &format!("price:{}", product.price),
            &format!("discountPercentage:{}", product.discount_percentage),
        ]);

        self.product_item = product;
    }
}

#[no_mangle]
fn main() {
    let mut data = Data::default();
    data.set_name("NewProduct");
    data.set_product(PRODUCT_JSON);

    println!("{:?}", data);
}
