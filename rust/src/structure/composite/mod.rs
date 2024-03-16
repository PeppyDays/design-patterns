pub mod file;
pub mod folder;
pub mod main;

pub trait Component {
    fn search(&self, keyword: &str);
}
