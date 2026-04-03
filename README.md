# 🔊 Operación: "¡No te apagues, maldito!" (Sound Inaudible)

¿Tienes un parlante que se cree inteligente y se apaga cada 5 minutos de silencio? ¿Estás harto de que la música se corte o que el primer segundo de un video no se escuche porque el parlante estaba "durmiendo"? **¡ESTA ES TU SOLUCIÓN!**

Esta aplicación en Go es básicamente un **despertador invisible** para tu parlante. Le da un "toquecito" cada 3 minutos para recordarle que su único propósito en la vida es sonar, no dormir la siesta.

## 🛠️ ¿Qué hace este invento?
- Lanza un pulso de **20Hz** (un sonido tan bajo que solo tu parlante y quizás un elefante pueden sentir).
- Dura solo **500ms** (medio segundo de pura rebeldía).
- El volumen es un susurro del **5%**.
- **Resultado:** Tu parlante detecta señal, se mantiene despierto, pero tú no escuchas absolutamente NADA. Es magia.

## 🚀 Cómo ponerlo en marcha (Antes de que el parlante se duerma de nuevo)

1. **La vía rápida:**
   Haz doble clic en `inaudible.exe` (ya lo dejé compilado para ti). Verás una ventanita de terminal que dice que está activo. ¡No la cierres!

2. **La vía "Soy Programador":**
   ```bash
   go run main.go
   ```

3. **Para que sea eterno:**
   Crea un acceso directo a `sound-inaudible.exe` y mételo en la carpeta de inicio de Windows (`shell:startup`). Así, cada vez que prendas la PC, el parlante recibirá sus "cachetadas virtuales" automáticas.

## 🛑 ¿Cómo lo detengo?
Como ahora el programa corre en **segundo plano** (sin ventana negra), para cerrarlo debes:
1. Abrir el **Administrador de Tareas** (`Ctrl + Shift + Esc`).
2. Buscar `sound-inaudible.exe`.
3. Dale a **Finalizar tarea**.

---

## 😤 Unas palabras finales del autor...

**¡SÍ, YA LO SÉ!** Sé que es una solución rústica. Sé que un parlante normal no debería hacer esto. Sé que hay mejores formas de gastar ciclos de CPU...

**¡PERO MI PARLANTE NO HACE CASO Y NO PIENSO BOTARLO AL BASURERO TODAVÍA!** 🗑️🚫

Si este pedacito de código salva tu salud mental y evita que un periférico perfectamente bueno termine en un vertedero, mi misión está cumplida. 🫡
