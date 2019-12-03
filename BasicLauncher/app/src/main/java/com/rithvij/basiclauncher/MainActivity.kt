package com.rithvij.basiclauncher

import android.app.WallpaperManager
import android.content.Intent
import android.os.Bundle
import android.util.Log
import android.view.GestureDetector
import android.view.MotionEvent
import android.view.View
import androidx.appcompat.app.AppCompatActivity
import kotlinx.android.synthetic.main.activity_main.*


class MainActivity : AppCompatActivity() {

    private val tag = "MainActivity"
    private lateinit var wallpaperManager: WallpaperManager
    private lateinit var mygestureDetector: GestureDetector

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        // Create a wallpaper manager instance
        wallpaperManager = WallpaperManager.getInstance(applicationContext)
        // Create a gesture detector instance
        mygestureDetector = GestureDetector(this@MainActivity, MyGestureDetector())


        if (wallpaperManager.wallpaperInfo != null){
            // If it is a livewallpaper
            Log.d(tag, wallpaperManager.wallpaperInfo.settingsActivity)
            Log.d(tag, wallpaperManager.wallpaperInfo.packageName)

            // clayout is the global constraint layout
            // It should be clickable. I did it in the xml
            val touchListener = View.OnTouchListener { _, event ->
                mygestureDetector.onTouchEvent(event)
            }
            clayout.setOnTouchListener(touchListener)

            button2.setOnClickListener {
                // To open current live wallpaper's activity
                // I answered this question on SO
                // https://stackoverflow.com/a/59156060/8608146
                val intent = Intent()
                    .setClassName(
                        wallpaperManager.wallpaperInfo.packageName,
                        wallpaperManager.wallpaperInfo.settingsActivity
                    )
                    .addFlags(Intent.FLAG_ACTIVITY_NEW_TASK)
                startActivity(intent)
            }
        } else {
            // remove button
            // (button2.parent as ViewGroup).removeView(button2)
            // or rename it
            button2.text = "Use a LiveWallpaper like Muzei to see how this works"
        }
    }

    inner class MyGestureDetector : GestureDetector.SimpleOnGestureListener() {
        init {
            Log.d(tag, "Init Gesture Detector")
        }

        override fun onDoubleTapEvent(e: MotionEvent?): Boolean {
            Log.d(tag, "double tap ${e!!.actionMasked} ${e.action}")
            handleGestures(e)
            return super.onDoubleTap(e)
        }

        override fun onSingleTapUp(e: MotionEvent?): Boolean {
            Log.d(tag, "single tap ${e!!.actionMasked} ${e.action}")
            handleGestures(e)
            return super.onSingleTapConfirmed(e)
        }

    }

    //    Lawnchair's implementation to forward gestures to the livewallpaper
    //    https://github.com/LawnchairLauncher/Lawnchair/blob/903a6398e81c1dda2991cc2292ae076e185c57b7/src/com/android/launcher3/Workspace.java#L1541
    fun handleGestures(event: MotionEvent?) {
        Log.d(tag, "gesture detection function")
        val position: Array<Int> = arrayOf(0, 0)
        clayout.getLocationOnScreen(position.toIntArray())
        // safe check
        event!!
        position[0] += event.x.toInt()
        position[1] += event.y.toInt()

        wallpaperManager.sendWallpaperCommand(
            clayout.windowToken,
            if (event.action == MotionEvent.ACTION_UP)
                WallpaperManager.COMMAND_TAP else
                WallpaperManager.COMMAND_SECONDARY_TAP,
            position[0], position[1], 0, null
        )
    }
}

