pub struct Account {
    name: String,
}

impl Account {
    pub fn new(name: String) -> Self {
        Self { name }
    }

    pub fn check(&self, name: &str) -> Result<(), String> {
        if &self.name != name {
            return Err("Account name is incorrect".into());
        }

        println!("Account verified");
        Ok(())
    }
}
