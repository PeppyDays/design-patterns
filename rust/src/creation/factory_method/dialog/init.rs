use super::gui::Dialog;
use super::html_gui::HtmlDialog;
use super::windows_gui::WindowsDialog;

pub fn initialize() -> &'static dyn Dialog {
    if cfg!(windows) {
        println!("-- Windows detected, creating Windows GUI --");
        &WindowsDialog
    } else {
        println!("-- No OS detected, creating the HTML GUI --");
        &HtmlDialog
    }
}
