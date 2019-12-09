package com.rithvij.basiclauncher

import android.app.PendingIntent
import android.app.WallpaperManager
import android.content.*
import android.content.pm.ActivityInfo
import android.content.pm.ApplicationInfo
import android.os.Build
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
    private lateinit var componentReceiver: ComponentReceiver

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        // Create a wallpaper manager instance
        wallpaperManager = WallpaperManager.getInstance(applicationContext)
        // Create a gesture detector instance
        mygestureDetector = GestureDetector(this@MainActivity, MyGestureDetector())
        componentReceiver = ComponentReceiver()

        updateUI()
    }

    override fun onResume() {
        super.onResume()
        val filter = IntentFilter()
        registerReceiver(componentReceiver, filter)
        updateUI()
    }

    // https://stackoverflow.com/a/18254831/8608146
    override fun onPause() {
        super.onPause()
        unregisterReceiver(componentReceiver)
    }

    private fun updateUI() {
        Log.d(tag, "UPDATING UI")
        // remove any listeners
        clayout.setOnTouchListener(null)
        button2.setOnClickListener(null)
        button.setOnClickListener(null)
        if (wallpaperManager.wallpaperInfo != null) {
            // If it is a livewallpaper
            var hasSettings = false
            // if it has a settings activity
            if (wallpaperManager.wallpaperInfo.settingsActivity != null) {
                Log.d(tag, wallpaperManager.wallpaperInfo.settingsActivity)
                hasSettings = true
            }
            // can be null
            // Log.d(
            //     tag, wallpaperManager.wallpaperInfo.loadDescription(packageManager).toString()
            // )

            // clayout is the global constraint layout
            // It should be clickable.
            clayout.isClickable = true
            clayout.isFocusable = true
            val touchListener = View.OnTouchListener { _, event ->
                mygestureDetector.onTouchEvent(event)
            }
            clayout.setOnTouchListener(touchListener)

            if (hasSettings) {
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
                button2.text = resources.getString(R.string.start_livewallpaper_activity)
            } else {
                button2.text = resources.getString(R.string.no_settings)
            }
        } else {
            // remove button
            // (button2.parent as ViewGroup).removeView(button2)
            // or rename it
            button2.text = resources.getString(R.string.no_livewallpaper)
        }

        button.setOnClickListener {
            val intent = Intent(WallpaperManager.ACTION_LIVE_WALLPAPER_CHOOSER)
            startActivity(intent)
        }

        wallpaper_btn.setOnClickListener {
            // https://stackoverflow.com/a/34314156/8608146
            val intent = Intent(Intent.ACTION_SET_WALLPAPER)

            // https://stackoverflow.com/a/18068122/8608146 this link is useful
            val providers = packageManager.queryIntentActivities(intent, 0)
            providers.forEach {
                Log.d(tag, it.activityInfo.packageName)
            }

            val debugApps = getDebugApps()
            debugApps.forEach {
                Log.d("$tag-debug", it.packageName)
            }

            // https://stackoverflow.com/a/19159291/8608146
            val receiver = Intent(this, ComponentReceiver::class.java)
            val pendingIntent =
                PendingIntent.getBroadcast(this, 0, receiver, PendingIntent.FLAG_UPDATE_CURRENT)
            // Create a chooser to prevent the user from checking don't ask again option
            val chooser = if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.LOLLIPOP_MR1) {
                Intent.createChooser(
                    intent,
                    resources.getString(R.string.set_wallpaper),
                    pendingIntent.intentSender
                )
            } else {
                Intent.createChooser(
                    intent,
                    resources.getString(R.string.set_wallpaper)
                )
            }
            startActivity(chooser)
        }

    }

    // https://stackoverflow.com/a/59242174/8608146
    private fun getDebugApps(): ArrayList<ActivityInfo> {
        val allIntent = Intent(Intent.ACTION_MAIN).addCategory(Intent.CATEGORY_LAUNCHER)
        val allApps = packageManager.queryIntentActivities(allIntent, 0)
        val debugApps = arrayListOf<ActivityInfo>()
        allApps.forEach {
            val appInfo = (packageManager.getApplicationInfo(
                it.activityInfo.packageName,
                0
            ))
            if (0 != appInfo.flags and ApplicationInfo.FLAG_DEBUGGABLE) {
                debugApps.add(it.activityInfo)
            }
        }

        return debugApps
    }

    // A receiver to get which one was chosen from the wallpaper chooser
    class ComponentReceiver : BroadcastReceiver() {
        override fun onReceive(p0: Context, p1: Intent) {
            // https://stackoverflow.com/questions/9583230/what-is-the-purpose-of-intentsender#comment72280489_34314156
            // EXTRA_CHOSEN_COMPONENT requires API 22
            if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.LOLLIPOP_MR1) {
                Log.d(
                    "MainActivity",
                    p1.extras!!.getParcelable<ComponentName>(Intent.EXTRA_CHOSEN_COMPONENT)!!.toString()
                )
            }
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

