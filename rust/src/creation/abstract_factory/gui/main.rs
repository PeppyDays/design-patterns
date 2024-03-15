// Abstract factory
pub trait Button {
    fn press(&self);
}

pub trait Checkbox {
    fn switch(&self);
}

pub trait GuiFactory {
    type B: Button;
    type C: Checkbox;

    fn create_button(&self) -> Self::B;
    fn create_checkbox(&self) -> Self::C;
}

// Windows factory
pub struct WindowsButton;

impl Button for WindowsButton {
    fn press(&self) {
        println!("Pressed windows button!");
    }
}

pub struct WindowsCheckbox;

impl Checkbox for WindowsCheckbox {
    fn switch(&self) {
        println!("Switched windows checkbox");
    }
}

pub struct WindowsGuiFactory;

impl GuiFactory for WindowsGuiFactory {
    type B = WindowsButton;
    type C = WindowsCheckbox;

    fn create_button(&self) -> Self::B {
        WindowsButton
    }

    fn create_checkbox(&self) -> Self::C {
        WindowsCheckbox
    }
}

// Mac factory
pub struct MacButton;

impl Button for MacButton {
    fn press(&self) {
        println!("Pressed mac button!");
    }
}

pub struct MacCheckbox;

impl Checkbox for MacCheckbox {
    fn switch(&self) {
        println!("Switched mac checkbox");
    }
}

pub struct MacGuiFactory;

impl GuiFactory for MacGuiFactory {
    type B = MacButton;
    type C = MacCheckbox;

    fn create_button(&self) -> Self::B {
        MacButton
    }

    fn create_checkbox(&self) -> Self::C {
        MacCheckbox
    }
}

pub fn render(factory: impl GuiFactory) {
    let button = factory.create_button();
    let checkbox = factory.create_checkbox();

    button.press();
    checkbox.switch();
}

fn main() {
    let factory = WindowsGuiFactory;

    render(factory)
}
